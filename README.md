<p align="center">
  <a href="https://github.com/kevincobain2000/instachart">
    <img alt="c2p" src="https://imgur.com/HC5FB7O.png" width="320">
  </a>
</p>
<p align="center">
  Chart as Image - service using <a href="https://github.com/wcharczuk/go-chart" target="_blank">go-chart</a>
 and <a href="https://github.com/vicanso/go-charts" target="_blank">go-charts.</a>
</p>

[![coverage](https://coveritup.app/embed/kevincobain2000/instachart?branch=master&type=coverage)](https://coveritup.app/kevincobain2000/instachart)

[![go-binary-size](https://coveritup.app/embed/kevincobain2000/instachart?branch=master&type=go-binary-size&color=pink)](https://coveritup.app/kevincobain2000/instachart)
[![unit-test-run-time](https://coveritup.app/embed/kevincobain2000/instachart?branch=master&type=unit-test-run-time)](https://coveritup.app/kevincobain2000/instachart)
[![build-time](https://coveritup.app/embed/kevincobain2000/instachart?branch=master&type=build-time&color=lightblue)](https://coveritup.app/kevincobain2000/instachart)
[![go-mod-dependencies](https://coveritup.app/embed/kevincobain2000/instachart?branch=master&type=go-mod-dependencies&color=indigo)](https://coveritup.app/kevincobain2000/instachart)
[![go-sec-issues](https://coveritup.app/embed/kevincobain2000/instachart?branch=master&type=go-sec-issues&color=green)](https://coveritup.app/kevincobain2000/instachart)

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
- [CHANGE LOG](#change-log)

# API

**URL** `https://instachart.coveritup.app`

| Query        | Required | Description | Default               |
| :----------- | :------- | :---------- | :-------------------- |
| `data`       | ◯        | JSON        |                       |
| `title`      |          | string      |                       |
| `subtitle`   |          | string      |                       |
| `theme`      |          | string      | `light` or `dark`     |
| `metric`     |          | string      |                       |
| `height`     |          | int         | 400                   |
| `width`      |          | int         | 600                   |
| `fill`       |          | boolean     | `false` for `/line`   |
| `horizontal` |          | boolean     | `false` for `/bar`    |



## `GET /line`

| `data`  | Required | Description      | Example                                                           |
| :------ | :------- | :--------------- | :---------------------------------------------------------------- |
| `x`     | ◯        | []Array (string) | `"x": [["2022-12-23","2022-12-24"], ["2022-12-23","2022-12-24"]]` |
| `y`     | ◯        | []Array (int)    | `"y": [[1,2],[3,4]]`                                              |
| `names` |          | Array            | `"names": ["Series A", "Series B"]`                               |

<details>
 <summary><b>REQUEST URL</b></summary>

```sh
https://instachart.coveritup.app/line?title=Line+Chart+Simple&subtitle=Sleeping+Hours&data={
    "x": [["Mon","Tue","Wed"]],
    "y": [[4,8,7], [10,20,24]],
    "names": ["Sleeping", "Awake"]
}
```

</details>

<br>


<p align="center">
  <a href='https://instachart.coveritup.app/line?title=Line+Chart&width=1024&height=620&subtitle=Sleeping+Hours&data={%20"x":%20[["Mon","Tue","Wed"]],%20"y":%20[[4,8,7],%20[10,20,24]],%20"names":%20["Sleeping",%20"Awake"]%20}'>
    <img alt="line chart" src='https://instachart.coveritup.app/line?title=Line+Chart&width=1024&height=620&subtitle=Sleeping+Hours&data={%20"x":%20[["Mon","Tue","Wed"]],%20"y":%20[[4,8,7],%20[10,20,24]],%20"names":%20["Sleeping",%20"Awake"]%20}'>
  </a>
</p>

<details>
 <summary><b>REQUEST URL</b></summary>

```sh
https://instachart.coveritup.app/line?title=Line+Chart+Simple&subtitle=Sleeping+Hours&data={
    "x": [["Mon","Tue","Wed"]],
    "y": [[4,8,7], [10,20,24]],
    "names": ["Sleeping", "Awake"]
}
```

</details>

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/line?title=Line+Chart+Area&width=1020&height=620&subtitle=Sleeping+Hours&fill=true&data={%20"x":%20[["Mon","Tue","Wed"]],%20"y":%20[[4,8,7],%20[10,20,24]],%20"names":%20["Sleeping",%20"Awake"]%20}'>
    <img alt="line area chart" src='https://instachart.coveritup.app/line?title=Line+Chart+Area&width=1020&height=620&subtitle=Sleeping+Hours&fill=true&data={%20"x":%20[["Mon","Tue","Wed"]],%20"y":%20[[4,8,7],%20[10,20,24]],%20"names":%20["Sleeping",%20"Awake"]%20}'>
  </a>
</p>


## `GET /bar`


| `data` | Required | Description    | Example                     |
| :----- | :------- | :------------- | :-------------------------- |
| `x`    | ◯        | Array (string) | `"x": ["Mon","Tue", "Wed"]` |
| `y`    | ◯        | []Array (int)  | `"y": [[1,2,3]]`            |

<details>
 <summary><b>REQUEST URL</b></summary>

```sh
https://instachart.coveritup.app/bar?title=Bar+Chart&subtitle=Sleeping+hours&data={
    "x": ["Monday", "Friday", "Sunday"],
    "y": [[8,2,14]]
}
```

</details>

<br>

<p align="center">
  <a href="https://instachart.coveritup.app/bar?title=Bar+Chart&subtitle=Sleeping+hours&data={%20%22x%22:%20[%22Monday%22,%20%22Friday%22,%20%22Sunday%22],%20%22y%22:%20[[8,%202%20,14]]%20}">
    <img alt="bar chart" src='https://instachart.coveritup.app/bar?title=Bar+Chart&subtitle=Sleeping+hours&data={%20%22x%22:%20[%22Monday%22,%20%22Friday%22,%20%22Sunday%22],%20%22y%22:%20[[8,%202,14]]%20}'>
  </a>
</p>

## `GET /donut`


| `data`   | Required | Description    | Example                     |
| :------- | :------- | :------------- | :-------------------------- |
| `names`  | ◯        | Array (string) | `"x": ["Mon","Tue", "Wed"]` |
| `values` | ◯        | Array (int)    | `"y": [1,2,3]`              |

<details>
 <summary><b>REQUEST URL</b></summary>

```sh
https://instachart.coveritup.app/donut?title=Donut+Chart&data={
    "names": ["Monday", "Friday", "Saturday", "Sunday"],
    "values": [4, 6 ,7, 9]
}
```

</details>

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/donut?title=Donut+Chart&width=350&height=350&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20}'>
    <img alt="donut chart" src='https://instachart.coveritup.app/donut?title=Donut+Chart&width=350&height=350&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20}'>
  </a>
</p>



## `GET /pie`

| `data`   | Required | Description    | Example                     |
| :------- | :------- | :------------- | :-------------------------- |
| `names`  | ◯        | Array (string) | `"x": ["Mon","Tue", "Wed"]` |
| `values` | ◯        | Array (int)    | `"y": [1,2,3]`              |


<details>
 <summary><b>REQUEST URL</b></summary>

```sh
https://instachart.coveritup.app/pie?title=Pie+Chart&subtitle=Sleeping+Hours&data={
    "names": ["Monday", "Friday", "Saturday", "Sunday"],
    "values": [4, 6 ,7, 9]
}
```

</details>

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/pie?title=Pie+Chart&subtitle=Sleeping+Hours&height=520&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20}'>
    <img alt="pie chart" src='https://instachart.coveritup.app/pie?title=Pie+Chart&subtitle=Sleeping+Hours&height=520&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20}'>
  </a>
</p>


## `GET /radar`

| `data`   | Required | Description    | Example                                 | Validation                         |
| :------- | :------- | :------------- | :-------------------------------------- | :--------------------------------- |
| `names`  | ◯        | Array (string) | `"names": ["Mon","Tue", "Wed"]`         | `>=3`                              |
| `values` | ◯        | []Array (int)  | `"values": [[1,2,3], [3,4,5]]`          | `count(names) == count(values[0])` |
| `labels` |          | Array (string) | `"labels": ["Work", "Relax", "Travel"]` |                                    |

<details>
 <summary><b>REQUEST URL</b></summary>

```sh
https://instachart.coveritup.app/radar?title=Radar+Chart&data={
    "names": ["Mon","Tue", "Wed", "Fri"],
    "labels": ["Work", "Relax", "Travel"],
    "values": [[1,2,3,4], [15,7,8,9], [15,17,5,7]]
}
```

</details>

<br>

<p align="center">
  <a href='https://instachart.coveritup.app/radar?title=Radar+Chart&data={%20%22names%22:%20[%22Mon%22,%22Tue%22,%20%22Wed%22,%20%22Fri%22],%20%22labels%22:%20[%22Work%22,%20%22Relax%22,%20%22Travel%22],%20%22values%22:%20[[1,2,3,4],%20[15,7,8,9],%20[15,17,5,7]]%20}'>
    <img alt="radar chart" src='https://instachart.coveritup.app/radar?title=Radar+Chart&&height=520&data={%20%22names%22:%20[%22Mon%22,%22Tue%22,%20%22Wed%22,%20%22Fri%22],%20%22labels%22:%20[%22Work%22,%20%22Relax%22,%20%22Travel%22],%20%22values%22:%20[[1,2,3,4],%20[15,7,8,9],%20[15,17,5,7]]%20}'>
  </a>
</p>


# CHANGE LOG

- **v1.0.0** - Initial release with `line`, `bar`, `donut`, `pie`, `radar` charts.
- **v1.0.1** - Bug fixes and code refactor.