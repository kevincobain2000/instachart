<p align="center">
  <a href="https://github.com/kevincobain2000/instachart">
    <img alt="c2p" src="https://imgur.com/HC5FB7O.png" width="320">
  </a>
</p>
<p align="center">
  Chart as Image - API using go-charts.
</p>

**Quick Setup:** Easy setup. Add charts as images in markdown, emails etc.

**Hassle Free:** Create charts as images with single URL.

**Supports:** Choose from Line, bar, donut, pie, scatter, bubble, stacked bar, etc.


# API

**URL** `https://instachart.coveritup.app`

- [API](#api)
  - [`GET /line`](#get-line)
    - [Single Line Series](#single-line-series)
    - [Multiple Line Series](#multiple-line-series)
    - [Continuous Series](#continuous-series)
  - [`GET /bar`](#get-bar)
    - [Bars](#bars)


## [`GET /line`](https://instachart.coveritup.app/line?title=Single+Line+series&x_label=dates&y_label=amount&data={%20%22x%22:%20[[%222022-12-23%22,%222022-12-24%22,%222023-12-25%22]],%20%22y%22:%20[[1,2,3]]%20})

| Query     | Required | Description | Example                                                    |
| :-------- | :------- | :---------- | :--------------------------------------------------------- |
| `data`    | ◯        | JSON        | `?data={ "x": [["2022-12-23","2023-12-25"]],"y": [[1,2]]}` |
| `title`   |          | string      |                                                            |
| `x_label` |          | string      |                                                            |
| `y_label` |          | string      |                                                            |
| `height`  |          | int         |                                                            |
| `width`   |          | int         |                                                            |


| `data`  | Required | Description | Example                                                           |
| :------ | :------- | :---------- | :---------------------------------------------------------------- |
| `x`     | ◯        | []Array     | `"x": [["2022-12-23","2022-12-24"], ["2022-12-23","2022-12-24"]]` |
| `y`     | ◯        | []Array     | `"y": [[1,2],[3,4]]`                                              |
| `names` |          | Array       | `"names": ["Series A", "Series B"]`                               |

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

![Line single series](https://instachart.coveritup.app/line?title=Single+Line+Series&x_label=dates&y_label=amount&data={%20%22x%22:%20[[%222022-12-23%22,%222022-12-24%22,%222023-12-25%22]],%20%22y%22:%20[[1,2,3]],%20%22names%22:%20[%22Series%20A%22]%20})

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

![Multi line series](https://instachart.coveritup.app/line?title=Multi+Line+Series&x_label=dates&y_label=amount&data={%20%22x%22:%20[[%222022-12-23%22,%222022-12-24%22,%222023-12-25%22],%20[%222022-12-23%22,%222022-12-28%22,%222023-12-30%22]],%20%22y%22:%20[[1,2,3],%20[1,5,10]],%20%22names%22:%20[%22Series%20A%22,%20%22Series%20B%22]%20})


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

![Multi line series](https://instachart.coveritup.app/line?title=Continuous+Series&x_label=No+of+people&y_label=amount&data={%20"x":%20[["10","20","30"],%20["10","20","30"],%20["10","20","30"]],%20"y":%20[[1,2,3],%20[10,20,30],%20[6,3,9]]%20})



## [`GET /bar`](https://instachart.coveritup.app/bar?title=Bar+Chart&y_label=Sleeping+hours&data={%20%22x%22:%20[%22Monday%22,%20%22Friday%22,%20%22Sunday%22],%20%22y%22:%20[8,%202%20,14]%20})



| Query     | Required | Description | Example                                              |
| :-------- | :------- | :---------- | :--------------------------------------------------- |
| `data`    | ◯        | JSON        | `?data={ "x": [["Mon","Tue","Wed"]],"y": [[1,2,3]]}` |
| `title`   |          | string      |                                                      |
| `y_label` |          | string      |                                                      |
| `height`  |          | int         |                                                      |
| `width`   |          | int         |                                                      |


| `data` | Required | Description | Example                        |
| :----- | :------- | :---------- | :----------------------------- |
| `x`    | ◯        | []Array     | `"x": [["Mon","Tue", "Wed"], ` |
| `y`    | ◯        | []Array     | `"y": [[1,2,3]]`               |


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

![Bar chart](https://instachart.coveritup.app/bar?title=Bar+Chart&y_label=Sleeping+hours&data={%20%22x%22:%20[%22Monday%22,%20%22Friday%22,%20%22Sunday%22],%20%22y%22:%20[8,%202%20,14]%20})