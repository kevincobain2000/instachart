<p align="center">
  <a href="https://github.com/kevincobain2000/instachart">
    <img alt="c2p" src="https://imgur.com/DyszMue.png" width="320">
  </a>
</p>
<p align="center">
  Chart as Image - service using <a href="https://github.com/wcharczuk/go-chart" target="_blank">go-chart</a>
 and <a href="https://github.com/vicanso/go-charts" target="_blank">go-charts.</a>
</p>
<p align="center">
    <a href="https://www.producthunt.com/posts/instachart?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-instachart" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=432315&theme=neutral" alt="InstaChart - Generate&#0032;charts&#0032;as&#0032;images | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>
</p>


![unit-test-run-time](https://coveritup.app/badge?org=kevincobain2000&repo=instachart&type=unit-test-run-time&branch=master)
![coverage](https://coveritup.app/badge?org=kevincobain2000&repo=instachart&type=coverage&branch=master)
![go-binary-size](https://coveritup.app/badge?org=kevincobain2000&repo=instachart&type=go-binary-size&branch=master)
![go-mod-dependencies](https://coveritup.app/badge?org=kevincobain2000&repo=instachart&type=go-mod-dependencies&branch=master)
![go-sec-issues](https://coveritup.app/badge?org=kevincobain2000&repo=instachart&type=go-sec-issues&branch=master)
![npm-install-time](https://coveritup.app/badge?org=kevincobain2000&repo=instachart&type=npm-install-time&branch=master)
![npm-modules-size](https://coveritup.app/badge?org=kevincobain2000&repo=instachart&type=npm-modules-size&branch=master)
![npm-build-time](https://coveritup.app/badge?org=kevincobain2000&repo=instachart&type=npm-build-time&branch=master)
![go-build-time](https://coveritup.app/badge?org=kevincobain2000&repo=instachart&type=go-build-time&branch=master)

![unit-test-run-time](https://coveritup.app/chart?org=kevincobain2000&repo=instachart&type=unit-test-run-time&output=svg&width=160&height=160&branch=master)
![coverage](https://coveritup.app/chart?org=kevincobain2000&repo=instachart&type=coverage&output=svg&width=160&height=160&branch=master)
![go-binary-size](https://coveritup.app/chart?org=kevincobain2000&repo=instachart&type=go-binary-size&output=svg&width=160&height=160&branch=master)
![go-mod-dependencies](https://coveritup.app/chart?org=kevincobain2000&repo=instachart&type=go-mod-dependencies&output=svg&width=160&height=160&branch=master)
![go-sec-issues](https://coveritup.app/chart?org=kevincobain2000&repo=instachart&type=go-sec-issues&output=svg&width=160&height=160&branch=master)
![npm-install-time](https://coveritup.app/chart?org=kevincobain2000&repo=instachart&type=npm-install-time&output=svg&width=160&height=160&branch=master)
![npm-modules-size](https://coveritup.app/chart?org=kevincobain2000&repo=instachart&type=npm-modules-size&output=svg&width=160&height=160&branch=master)
![npm-build-time](https://coveritup.app/chart?org=kevincobain2000&repo=instachart&type=npm-build-time&output=svg&width=160&height=160&branch=master)
![go-build-time](https://coveritup.app/chart?org=kevincobain2000&repo=instachart&type=go-build-time&output=svg&width=160&height=160&branch=master)


**Quick Setup:** Easy setup. Add charts as images in markdown, emails etc.

**Hassle Free:** Create charts as images with single URL.

**Supports:** Choose from Line, bar, donut, pie, scatter, bubble, stacked bar, etc.

**Lightweight:** Written in Go. No external dependencies. Download and self-host binary.

**API Documentation** Restful API with `image/png` response.

- [API](#api)
  - [`GET /line`](#get-line)
  - [`GET /bar`](#get-bar)
  - [`GET /donut`](#get-donut)
  - [`GET /pie`](#get-pie)
  - [`GET /radar`](#get-radar)
  - [`GET /funnel`](#get-funnel)
  - [`GET /table`](#get-table)
- [Self Hosting](#self-hosting)
- [CHANGE LOG](#change-log)

# API

**URL** `https://instachart.coveritup.app`

| Common Queries | Required | Type   | Default | Description       |
| :------------- | :------- | :----- | :------ | :---------------- |
| `data`         | ◯        | string |         | JSON Format       |
| `title`        |          | string |         |                   |
| `subtitle`     |          | string |         |                   |
| `theme`        |          | string | `light` | `light` or `dark` |
| `metric`       |          | string |         |                   |
| `height`       |          | int    | 768     |                   |
| `width`        |          | int    | 1024    |                   |
| `output`       |          | string | `png`   | `png` or `svg`    |



## `GET /line`

| Query  | Required | Type   | Default | Description        |
| :----- | :------- | :----- | :------ | :----------------- |
| `line` |          | string | `fill`  | `fill` or `nofill` |


| `data`  | Required | Type     |
| :------ | :------- | :------- |
| `x`     | ◯        | []string |
| `y`     | ◯        | [][]int  |
| `names` |          | []string |


```sh
https://instachart.coveritup.app/line?title=Line+Chart
&subtitle=Sleeping+Hours
&data={
    "x": ["Mon","Tue","Wed"],
    "y": [[4,8,7], [10,20,24]],
    "names": ["Sleeping", "Awake"]
}
```

<br>


<p align="center">
  <a href='https://instachart.coveritup.app/line?title=Line+Chart&width=1024&height=620&subtitle=Sleeping+Hours&data={%20"x":%20["Mon","Tue","Wed"],%20"y":%20[[4,8,7],%20[10,20,24]],%20"names":%20["Sleeping",%20"Awake"]%20}'>
    <img alt="line chart" src='https://instachart.coveritup.app/line?title=Line+Chart&width=1024&height=620&subtitle=Sleeping+Hours&data={%20"x":%20["Mon","Tue","Wed"],%20"y":%20[[4,8,7],%20[10,20,24]],%20"names":%20["Sleeping",%20"Awake"]%20}'>
  </a>
</p>

```sh
https://instachart.coveritup.app/line?title=Area+Chart
&subtitle=Sleeping+Hours
&fill=true
&data={
    "x": ["Mon","Tue","Wed"],
    "y": [[4,8,7], [10,20,24]],
    "names": ["Sleeping", "Awake"]
}
```

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/line?title=Line+Chart+Area&width=1020&height=620&subtitle=Sleeping+Hours&fill=true&data={%20"x":%20["Mon","Tue","Wed"],%20"y":%20[[4,8,7],%20[10,20,24]],%20"names":%20["Sleeping",%20"Awake"]%20}'>
    <img alt="line area chart" src='https://instachart.coveritup.app/line?title=Line+Chart+Area&width=1020&height=620&subtitle=Sleeping+Hours&fill=true&data={%20"x":%20["Mon","Tue","Wed"],%20"y":%20[[4,8,7],%20[10,20,24]],%20"names":%20["Sleeping",%20"Awake"]%20}'>
  </a>
</p>


## `GET /bar`

Extra params for `bar` chart.

| Query     | Required | Type   | Default    | Validation                 |
| :-------- | :------- | :----- | :--------- | :------------------------- |
| `style`   |          | string | `vertical` | `vertical` or `horizontal` |
| `zmetric` |          | string | `vertical` | when `z` data is provided  |


| `data` | Required | Type     | Description                                |
| :----- | :------- | :------- | :----------------------------------------- |
| `x`    | ◯        | []string |                                            |
| `y`    | ◯        | [][]int  |                                            |
| `z`    |          | []string | Style will be `vertical` with `z` is given |

```sh
https://instachart.coveritup.app/bar?title=Bar+Chart
&subtitle=Sleeping+hours
&data={
    "x": ["Monday", "Friday", "Sunday"],
    "y": [[8,2,14]],
    "names": ["Sleeping", "Awake"]
}
```

<br>

<p align="center">
  <a href="https://instachart.coveritup.app/bar?title=Bar+Chart&subtitle=Sleeping+hours&data={%20%22x%22:%20[%22Monday%22,%20%22Friday%22,%20%22Sunday%22],%20%22y%22:%20[[8,%202%20,14]]%20}">
    <img alt="bar chart" src='https://instachart.coveritup.app/bar?title=Bar+Chart&subtitle=Sleeping+hours&data={%20%22x%22:%20[%22Monday%22,%20%22Friday%22,%20%22Sunday%22],%20%22y%22:%20[[8,%202,14]]%20}'>
  </a>
</p>

<br>

```sh
https://instachart.coveritup.app/bar?title=Bar+Chart
&subtitle=Sleeping+hours
&metric=hours
&zmetric=days
&theme=dark
&data={
    "x": ["Monday", "Friday", "Sunday"],
    "y": [[8,2,14]],
    "z": [[9,3,15]],
    "names": ["Sleeping", "Awake"]
}
```

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/bar?title=Medidation+times&width=860&theme=dark&subtitle=Yoga&theme=light&metric=mins&width=1024&height=512&data={%22x%22:[%22Sunday%22,%22Monday%22,%22Monday%22,%22Monday%22,%22Tuesday%22,%22Tuesday%22,%22Tuesday%22,%22Wednesday%22,%22Wednesday%22,%22Wednesday%22,%22Wednesday%22,%22Wednesday%22,%22Thursday%22,%22Thursday%22,%22Thursday%22,%22Thursday%22,%22Friday%22,%22Friday%22,%22Friday%22,%22Saturday%22],%22y%22:[[19,19,72,66,80,86,148,145,152,126,107,150,133,106,204,105,106,103,18,104]],%22z%22:[[20,20,73,67,81,87,149,146,153,127,108,151,134,107,205,106,107,104,19,105]]}'>
    <img alt="bar chart" src='https://instachart.coveritup.app/bar?title=Medidation+times&width=860&theme=dark&subtitle=Yoga&theme=light&metric=mins&width=1024&height=512&data={%22x%22:[%22Sunday%22,%22Monday%22,%22Monday%22,%22Monday%22,%22Tuesday%22,%22Tuesday%22,%22Tuesday%22,%22Wednesday%22,%22Wednesday%22,%22Wednesday%22,%22Wednesday%22,%22Wednesday%22,%22Thursday%22,%22Thursday%22,%22Thursday%22,%22Thursday%22,%22Friday%22,%22Friday%22,%22Friday%22,%22Saturday%22],%22y%22:[[19,19,72,66,80,86,148,145,152,126,107,150,133,106,204,105,106,103,18,104]],%22z%22:[[20,20,73,67,81,87,149,146,153,127,108,151,134,107,205,106,107,104,19,105]]}'>
  </a>
</p>

## `GET /donut`


| `data`   | Required | Type     |
| :------- | :------- | :------- |
| `names`  | ◯        | []string |
| `values` | ◯        | []int    |

```sh
https://instachart.coveritup.app/donut?title=Donut+Chart
&data={
    "names": ["Monday", "Friday", "Saturday", "Sunday"],
    "values": [4, 6 ,7, 9]
}
```

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/donut?title=Donut+Chart&width=450&height=350&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20}'>
    <img alt="donut chart" src='https://instachart.coveritup.app/donut?title=Donut+Chart&width=450&height=350&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20}'>
  </a>
</p>



## `GET /pie`

| `data`   | Required | Type     |
| :------- | :------- | :------- |
| `names`  | ◯        | []string |
| `values` | ◯        | []int    |


```sh
https://instachart.coveritup.app/pie?title=Pie+Chart
&subtitle=Sleeping+Hours
&data={
    "names": ["Monday", "Friday", "Saturday", "Sunday"],
    "values": [4, 6 ,7, 9]
}
```

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/pie?title=Pie+Chart&subtitle=Sleeping+Hours&height=450&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20}'>
    <img alt="pie chart" src='https://instachart.coveritup.app/pie?title=Pie+Chart&subtitle=Sleeping+Hours&height=450&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20}'>
  </a>
</p>


## `GET /radar`

| `data`   | Required | Type     | Validation                         |
| :------- | :------- | :------- | :--------------------------------- |
| `names`  | ◯        | []string | `>=3`                              |
| `values` | ◯        | [][]int  | `count(names) == count(values[0])` |
| `labels` |          | []string |                                    |

```sh
https://instachart.coveritup.app/radar?title=Radar+Chart
&data={
    "names": ["Mon","Tue", "Wed", "Fri"],
    "labels": ["Work", "Relax", "Travel"],
    "values": [[1,2,3,4], [15,7,8,9], [15,17,5,7]]
}
```

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/radar?title=Radar+Chart&theme=dark&data={%20%22names%22:%20[%22Mon%22,%22Tue%22,%20%22Wed%22,%20%22Fri%22],%20%22labels%22:%20[%22Work%22,%20%22Relax%22,%20%22Travel%22],%20%22values%22:%20[[1,2,3,4],%20[15,7,8,9],%20[15,17,5,7]]%20}'>
    <img alt="radar chart" src='https://instachart.coveritup.app/radar?title=Radar+Chart&theme=dark&height=520&data={%20%22names%22:%20[%22Mon%22,%22Tue%22,%20%22Wed%22,%20%22Fri%22],%20%22labels%22:%20[%22Work%22,%20%22Relax%22,%20%22Travel%22],%20%22values%22:%20[[1,2,3,4],%20[15,7,8,9],%20[15,17,5,7]]%20}'>
  </a>
</p>

## `GET /funnel`

| `data`   | Required | Type     |
| :------- | :------- | :------- |
| `names`  | ◯        | []string |
| `values` | ◯        | []int    |

```sh
https://instachart.coveritup.app/funnel?title=Radar+Chart
&data={
    "names": ["Mon","Tue", "Thu", "Fri", "Sat", "Sun"],
    "values": [100,80,60,40,20,10]
}
```

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/funnel?title=Radar+Chart&width=620&data={%20%22names%22:%20[%22Mon%22,%22Tue%22,%20%22Thu%22,%20%22Fri%22,%20%22Sat%22,%20%22Sun%22],%20%22values%22:%20[100,80,60,40,20,10]%20}'>
    <img alt="funnel chart" src='https://instachart.coveritup.app/funnel?title=Radar+Chart&width=620&data={%20%22names%22:%20[%22Mon%22,%22Tue%22,%20%22Thu%22,%20%22Fri%22,%20%22Sat%22,%20%22Sun%22],%20%22values%22:%20[100,80,60,40,20,10]%20}'>
  </a>
</p>

## `GET /table`

| `data`   | Required | Type       | Description  |
| :------- | :------- | :--------- | :----------- |
| `names`  | ◯        | []string   | `aka header` |
| `values` | ◯        | [][]string | `aka rows`   |

```sh
https://instachart.coveritup.app/table?title=Table+Chart
&data={
    "names": ["Branch","Code Coverage", "Quality"],
    "values": [["master","80","90"], ["develop","70","79"]]
}
```

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/table?data={%20"names":%20["Branch","Code%20Coverage",%20"Quality"],%20"values":%20[["master","80","90"],%20["develop","70","79"]]%20}'>
    <img alt="table" src='https://instachart.coveritup.app/table?data={%20"names":%20["Branch","Code%20Coverage",%20"Quality"],%20"values":%20[["master","80","90"],%20["develop","70","79"]]%20}'>
  </a>
</p>


# Self Hosting

```sh
curl -sLk https://raw.githubusercontent.com/kevincobain2000/instachart/master/install.sh | sh
## Will run server on localhost:3001
# For more options see --help
./instachart
```

# CHANGE LOG

- **v1.0.0** - Initial release with `line`, `bar`, `donut`, `pie`, `radar`.
- **v1.0.1** - Bug fixes and code refactor.
- **v1.0.2** - `funnel` charts and `table` as image.
- **v1.0.3** - Fixes a few padding issues.
- **v1.0.4** - Adds `/` for frontend. And Installation instructions for self-hosting.
- **v1.0.5** - Frontend page with examples.
- **v1.0.6** - Code improvements and bug fixes.
- **v1.0.7** - Adds trends for `bar` charts.
- **v1.0.8** - Adds support for `svg` charts.
- **v1.0.9** - Validations improvements.
- **v1.0.11** - No LP on release version.
- **v1.0.14** - Support mini
- **v1.0.15** - Do not show grids on mini charts
