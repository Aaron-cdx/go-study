# prometheus语法学习

## prometheus查询
Prometheus提供了叫做PromQL(Prometheus Query Language)的功能型查询语言，可以用户选择和聚合一段时间序列的实时数据。这些表达式的结果不仅可以作为图像展示，
在Prometheus的表达式浏览器中查看表格数据，也可以通过HTTP API由外部系统使用。

### 表达式语言的数据类型
在Prometheus的表达式语言中，表达式或者子表达式可以计算为一下四种类型之一：
* 即时向量：一组时间序列，每个时间序列包含一个样本，所有的时间序列都共享相同的时间戳
* 范围向量：一组时间序列，包含每个时间序列跟随时间变化的数据点范围
* 标量(Scalar)：一个简单的数字浮点值
* 字符串(String)：~~一个简单的字符串值，目前未使用~~
根据用例(例如，在绘制和现实表达式的输出的时候)，作为用户指定表达式的结果，只要其中一些类型是合法的。例如，**返回即时向量的表达式是唯一可以直接绘制的类型**。

#### 时间序列选择器
**即时向量选择器**
即时向量选择器允许在给定的时间戳(即时)选择一组的时间序列和单个样本值：在最简单的形式中，只指定一个度量名称，这会产生一个包含所有具有此度量名称的时间序列元素的即时向量。
示例表示选择具有`http_requests_total`指标名称的所有时间序列。
> http_requests_total

可以通过花括号{}中附加逗号分隔的标签匹配器列表来进一步过滤这些时间序列。
下面示例表示选择具有http_requests_total指标名称且作业标签设置为prometheus，组标签设置为canary的时间序列
> http_requests_total{job="prometheus",group="canary"}

也可以否定匹配标签值，或将标签值与正则表达式进行匹配，存在一下标签匹配运算符：
* `=`: 选择与提供的字符串完全相同的标签
* `!=`: 选择与提供的字符串不相同的标签
* `=~`: 选择与提供的字符串正则匹配的标签
* `!~`: 选择与提供的字符串正则不匹配的标签

下面示例表示选择所有在staging/testing/development环境的http_requests_total且HTTP的方法不是GET的时间序列
 > http_requests_total{environment=~"staging|testing|development",method!="GET"}

匹配空标签值的标签匹配器还会选择根本没有设置特定标签的所有时间序列，正则表达式匹配是完全锚定的，同一个标签名称可以有多个匹配器。
向量选择器必须制定一个名称或至少一个与空字符串不匹配的标签匹配器，以下表达式是非法的：
> {job=~".*"} # Bad!

相反，这些表达式是有效的，因为它们都有一个不匹配空标签值的选择器
> {job=~".+"} # Good!
> {job=~".*", method="get"} # Good!

标签匹配器也可以通过内部的`__name__`标签来应用于度量名称，例如，表达式http_requests_total等价于`{__name__="http_requests_total"}`.
也可以使用=(!=,=~,!~)以外的匹配器。下面的表达式选择了所有名称以job开头的度量：
> {__name__=~"job:.*"}

指标名称不能够是关键字，如果关键字建议使用上面的`__name__`来指定具体的名称，然后查询数据。

**范围向量选择器**
范围向量字面量的工作方式与即时向量字面量类似，不同之处在于它们从当前时刻选择了一系列样本，从语法上说，持续时间在向量选择器的末尾附加在方括号中[]，
以指定应该为每个结果范围向量元素提取多远的时间值。
在此示例中，我们为所有具有指标名称http_requests_total和设置为prometheus的作业标签的时间序列选择我们过去5分钟内记录的所有值：
> http_requests_total{job="prometheus"}[5m]

**持续时间**
持续时间可以是一些确切的数字，后面跟上以下的时间单位：
* ms -> milliseconds
* s -> seconds
* m -> minutes
* h -> hours
* d -> days - assuming a day has always 24h
* w -> weeks - assuming a week has always 7d
* y -> years - assuming a year has always 365d

**偏移修饰符(Offset modifier)**
偏移修饰符允许更改查询中单个瞬时和范围向量的时间偏移。
例如，以下表达式返回相对于查询评估时间过去5分钟的http_requests_total的值：
> http_requests_total offset 5m

注意: 偏移修饰符需要始终跟随在选择器后面，否则会不正确。
> sum(http_requests_total{method="GET"} offset 5m)  # Good!
> sum(http_requests_total{method="GET"}) offset 5m  # Bad!

这个同样适用于范围向量，这将返回http_requests_total一周前的5分钟速率：
> rate(http_requests_total[5m] offset 1w)

为了与时间上的时间迁移进行比较，可以指定负偏移量(但这需要启用设置`--enable-feature=promql-negative-offset`标签)
> rate(http_requests_total[5m] offset -1w)

**修饰符**
@修饰符允许更改查询中单个即时和范围向量的评估时间，提供给@修饰符的时间是一个unix时间戳并用浮点文字描述。
> http_requests_total @ 1609746000 # 表示当前时间戳的数据

@修饰符也需要跟在选择器后面
> sum(http_requests_total{method="GET"} @ 1609746000) # GOOD
> sum(http_requests_total{method="GET"}) @ 1609746000 # Invalid

同样适用于范围向量：
> range(http_requests_total[5m] @ 1609746000) # GOOD

@修饰符支持int64限制内的所有的浮点文字表示，可以与偏移修饰符一起使用，其中偏移是相对@修饰符时间应用的，而不管哪个修饰符被先写入，下面两个查询将展示同样的结果
> # offset after @
> http_requests_total @ 1609746000 offset 5m
> # offset before @
> http_requests_total offset 5m @ 1609746000

默认情况下是禁用@修饰符的，可以通过`--enable-feature=promql-at-modifier`

除此之外，start()和end()同样可以作为@修饰符来表示特定的值
对于范围查询，它们分别解析为范围查询的开始和结束，并在所有步骤中保持不变
对于即时查询，start()和end()都解析为评估时间
> http_requests_total @ start()
> rate(http_requests_total)

**子查询**
子查询允许你对给定的范围和分辨率运行即时查询，子查询的结果是一个范围向量。
**操作符**
**函数**
**注释**
使用#号可以写注释
> # this is a prometheus comment.

**避免慢查询和负载**
如果一个查询需要对了大量的数据进行操作，则绘制它可能会超时或者使浏览器或者服务器过载。因此在对位置数据构建查询的时候，始终在Prometheus
表达式浏览器的表格视图中开始构建查询，直到结果集看起来合理(最多成百上千个时间序列)，只有当你充分过滤或者聚合数据时，才切换到图形模式。
如果表达式仍然需要很长的时间来临时绘制图形，请通过记录规则预先记录它。
这与prometheus的查询语言尤其相关，其中像api_http_requests_total这样的裸标签名称选择器可以扩展到具有不同标签的数千个时间序列。还有记住，
即使输出只是少量的时间序列，聚合多个时间序列的表达式也会在服务器上产生负载。这类似于在关系型数据库中对列的所有值求和会很慢，即使输出值只是一个数字。


















