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


# API

**URL** `https://instachart.coveritup.app`

- [API](#api)
  - [`GET /line`](#get-line)
    - [Single Line Series](#single-line-series)
    - [Multiple Line Series](#multiple-line-series)
    - [Continuous Series](#continuous-series)
  - [`GET /bar`](#get-bar)
    - [Bars](#bars)
  - [`GET /donut`](#get-donut)
    - [Donuts](#donuts)
  - [`GET /pie`](#get-pie)
    - [Pies](#pies)
  - [`GET /radar`](#get-radar)
    - [Radars](#radars)


## [`GET /line`](https://instachart.coveritup.app/line?title=Single+Line+series&x_label=dates&y_label=amount&data={%20%22x%22:%20[[%222022-12-23%22,%222022-12-24%22,%222023-12-25%22]],%20%22y%22:%20[[1,2,3]]%20})

| Query     | Required | Description | Example                                                    |
| :-------- | :------- | :---------- | :--------------------------------------------------------- |
| `data`    | ◯        | JSON        | `?data={ "x": [["2022-12-23","2023-12-25"]],"y": [[1,2]]}` |
| `title`   |          | string      |                                                            |
| `fill`    |          | boolean     |                                                            |
| `x_label` |          | string      |                                                            |
| `y_label` |          | string      |                                                            |
| `height`  |          | int         |                                                            |
| `width`   |          | int         |                                                            |
| `color`   |          | string      | `blue`,`red`,`green`,`yellow`,`cyan`,`orange`              |


| `data`  | Required | Description      | Example                                                           |
| :------ | :------- | :--------------- | :---------------------------------------------------------------- |
| `x`     | ◯        | []Array (string) | `"x": [["2022-12-23","2022-12-24"], ["2022-12-23","2022-12-24"]]` |
| `y`     | ◯        | []Array (int)    | `"y": [[1,2],[3,4]]`                                              |
| `names` |          | Array            | `"names": ["Series A", "Series B"]`                               |

### Single Line Series

<details>
 <summary><b>REQUEST URL</b></summary>

```sh
https://instachart.coveritup.app/line?title=Single+Line+Series&x_label=dates&y_label=amount&data={
    "x": [["2022-12-23","2022-12-24","2023-12-25"]],
    "y": [[1,2,3]]
}
```
</details>

<br>

![Line single series](https://instachart.coveritup.app/line?title=Single+Line+Series&width=460&height=310&x_label=dates&y_label=amount&data={%20%22x%22:%20[[%222022-12-23%22,%222022-12-24%22,%222023-12-25%22]],%20%22y%22:%20[[1,2,3]],%20%22names%22:%20[%22Series%20A%22]%20})
![Line single series](https://instachart.coveritup.app/line?title=Single+Line+Series&width=460&height=310&&fill=true&x_label=dates&y_label=amount&data={%20%22x%22:%20[[%222022-12-23%22,%222022-12-24%22,%222023-12-25%22]],%20%22y%22:%20[[1,2,3]],%20%22names%22:%20[%22Series%20A%22]%20})

### Multiple Line Series

<details>
 <summary><b>REQUEST URL</b></summary>

```sh
https://instachart.coveritup.app/line?title=Multi+Line+Series&x_label=dates&y_label=amount&data={
    "x": [["2022-12-23","2022-12-24","2023-12-25"], ["2022-12-23","2022-12-28","2023-12-30"]],
    "y": [[1,2,3], [1,5,10]],
    "names": ["Series A", "Series B"]
}
```
</details>

<br>

![Multi line series](https://instachart.coveritup.app/line?title=Multi+Line+Series&width=460&height=310&x_label=dates&y_label=amount&data={%20%22x%22:%20[[%222022-12-23%22,%222022-12-24%22,%222023-12-25%22],%20[%222022-12-23%22,%222022-12-28%22,%222023-12-30%22]],%20%22y%22:%20[[1,2,3],%20[1,5,10]],%20%22names%22:%20[%22Series%20A%22,%20%22Series%20B%22]%20})
![Multi line series](https://instachart.coveritup.app/line?title=Multi+Line+Series&width=460&height=310&fill=true&x_label=dates&y_label=amount&data={%20%22x%22:%20[[%222022-12-23%22,%222022-12-24%22,%222023-12-25%22],%20[%222022-12-23%22,%222022-12-28%22,%222023-12-30%22]],%20%22y%22:%20[[1,2,3],%20[1,5,10]],%20%22names%22:%20[%22Series%20A%22,%20%22Series%20B%22]%20})


### Continuous Series

<details>
 <summary><b>REQUEST URL</b></summary>

```sh
https://instachart.coveritup.app/line?title=Continuous+Series&x_label=No+of+people&y_label=amount&data={
    "x": [["10","20","30"], ["10","20","30"], ["10","20","30"]],
    "y": [[1,2,3], [10,20,30], [6,3,9]],
    "names": ["Series A", "Series B", "Series C"]
}
```

</details>

<br>

![Continuous line series](https://instachart.coveritup.app/line?title=Continuous+Series&x_label=No+of+people&y_label=amount&width=460&height=310&data={%20"x":%20[["10","20","30"],%20["10","20","30"],%20["10","20","30"]],%20"y":%20[[1,2,3],%20[10,20,30],%20[6,3,9]]%20})
![Continuous line series](https://instachart.coveritup.app/line?title=Continuous+Series&x_label=No+of+people&y_label=amount&width=460&height=310&fill=true&data={%20"x":%20[["10","20","30"],%20["10","20","30"],%20["10","20","30"]],%20"y":%20[[1,2,3],%20[10,20,30],%20[6,3,9]]%20})



## [`GET /bar`](https://instachart.coveritup.app/bar?title=Bar+Chart&y_label=Sleeping+hours&data={%20%22x%22:%20[%22Monday%22,%20%22Friday%22,%20%22Sunday%22],%20%22y%22:%20[8,%202%20,14]%20})



| Query        | Required | Description | Example                                              |
| :----------- | :------- | :---------- | :--------------------------------------------------- |
| `data`       | ◯        | JSON        | `?data={ "x": [["Mon","Tue","Wed"]],"y": [[1,2,3]]}` |
| `title`      |          | string      |                                                      |
| `y_label`    |          | string      |                                                      |
| `base_value` |          | int         |                                                      |
| `height`     |          | int         |                                                      |
| `width`      |          | int         |                                                      |


| `data` | Required | Description      | Example                      |
| :----- | :------- | :--------------- | :--------------------------- |
| `x`    | ◯        | []Array (string) | `"x": [["Mon","Tue", "Wed"]` |
| `y`    | ◯        | []Array (int)    | `"y": [[1,2,3]]`             |


### Bars

<details>
 <summary><b>REQUEST URL</b></summary>

```sh
https://instachart.coveritup.app/bar?title=Bar+Chart&y_label=Sleeping+hours&data={
    "x": ["Monday", "Friday", "Sunday"],
    "y": [8, 2 ,14]
}
```

</details>

<br>

![Bar chart](https://instachart.coveritup.app/bar?title=Bar+Chart&y_label=Sleeping+hours&data={%20%22x%22:%20[%22Mon%22,%20%22Tue%22,%20%22Wed%22,%20%22Thu%22,%20%22Fri%22,%20%22Sat%22,%20%22Sunday%22],%20%22y%22:%20[2,%203,%206%20,14,%2020,%2021,%2040]%20})

## [`GET /donut`](https://instachart.coveritup.app/donut?title=Donut+Chart&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20})


| Query    | Required | Description | Example                                                   |
| :------- | :------- | :---------- | :-------------------------------------------------------- |
| `data`   | ◯        | JSON        | `?data={ "names": ["Mon","Tue","Wed"],"values": [1,2,3]}` |
| `title`  |          | string      |                                                           |
| `height` |          | int         |                                                           |
| `width`  |          | int         |                                                           |


| `data`   | Required | Description    | Example                     |
| :------- | :------- | :------------- | :-------------------------- |
| `names`  | ◯        | Array (string) | `"x": ["Mon","Tue", "Wed"]` |
| `values` | ◯        | Array (int)    | `"y": [1,2,3]`              |


### Donuts

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

![Donut chart](https://instachart.coveritup.app/donut?title=Donut+Chart&width=420&height=420&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20})


## [`GET /pie`](https://instachart.coveritup.app/pie?title=Pie+Chart&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20})


| Query      | Required | Description | Example                                                   |
| :--------- | :------- | :---------- | :-------------------------------------------------------- |
| `data`     | ◯        | JSON        | `?data={ "names": ["Mon","Tue","Wed"],"values": [1,2,3]}` |
| `title`    |          | string      |                                                           |
| `subtitle` |          | string      |                                                           |
| `height`   |          | int         |                                                           |
| `width`    |          | int         |                                                           |


| `data`   | Required | Description    | Example                     |
| :------- | :------- | :------------- | :-------------------------- |
| `names`  | ◯        | Array (string) | `"x": ["Mon","Tue", "Wed"]` |
| `values` | ◯        | Array (int)    | `"y": [1,2,3]`              |


### Pies

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

![Donut chart](https://instachart.coveritup.app/pie?title=Pie+Chart&subtitle=Sleeping+Hours&width=312&height=312&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20})


## [`GET /radar`](https://instachart.coveritup.app/radar?title=Pie+Chart&subtitle=Sleeping+Hours&data={%20"names":%20["Monday",%20"Friday",%20"Saturday",%20"Sunday"],%20"values":%20[4,%206%20,7,%209]%20})


| Query    | Required | Description | Example                                                                                  |
| :------- | :------- | :---------- | :--------------------------------------------------------------------------------------- |
| `data`   | ◯        | JSON        | `data={"names": ["Mon","Tue", "Wed"], "labels": ["Work", "Relax"], "values": [[1,2,4]]}` |
| `title`  |          | string      |                                                                                          |
| `height` |          | int         |                                                                                          |
| `width`  |          | int         |                                                                                          |


| `data`   | Required | Description    | Example                                 | Valid                              |
| :------- | :------- | :------------- | :-------------------------------------- | :--------------------------------- |
| `names`  | ◯        | Array (string) | `"names": ["Mon","Tue", "Wed"]`         | `>=3`                              |
| `values` | ◯        | []Array (int)  | `"values": [[1,2,3], [3,4,5]]`          | `count(names) == count(values[0])` |
| `labels` |          | Array (string) | `"labels": ["Work", "Relax", "Travel"]` |                                    |


### Radars

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

![Radar chart](https://instachart.coveritup.app/radar?title=Radar+Chart&data={%20%22names%22:%20[%22Mon%22,%22Tue%22,%20%22Wed%22,%20%22Fri%22],%20%22labels%22:%20[%22Work%22,%20%22Relax%22,%20%22Travel%22],%20%22values%22:%20[[1,2,3,4],%20[15,7,8,9],%20[15,17,5,7]]%20})

