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
| `grid` |          | string | `show`  | `show` or `hide`   |


| `data`  | Required | Type     |
| :------ | :------- | :------- |
| `x`     | ◯        | []string |
| `y`     | ◯        | [][]int  |
| `names` |          | []string |


```sh
/line?title=Line+Chart
&subtitle=Sleeping+Hours
&data={
    "x": ["Mon","Tue","Wed"],
    "y": [[4,8,7], [10,20,24]],
    "names": ["Sleeping", "Awake"]
}
```


```sh
/line?title=Area+Chart
&subtitle=Sleeping+Hours
&fill=true
&data={
    "x": ["Mon","Tue","Wed"],
    "y": [[4,8,7], [10,20,24]],
    "names": ["Sleeping", "Awake"]
}
```

<br>


## `GET /bar`

Extra params for `bar` chart.

| Query     | Required | Type   | Default    | Validation                 |
| :-------- | :------- | :----- | :--------- | :------------------------- |
| `style`   |          | string | `vertical` | `vertical` or `horizontal` |
| `grid`    |          | string | `show`     | `show` or `hide`           |


| `data` | Required | Type     | Description                                |
| :----- | :------- | :------- | :----------------------------------------- |
| `x`    | ◯        | []string |                                            |
| `y`    | ◯        | [][]int  |                                            |
| `z`    |          | []string | Style will be `vertical` with `z` is given |

```sh
/bar?title=Bar+Chart
&subtitle=Sleeping+hours
&data={
    "x": ["Monday", "Friday", "Sunday"],
    "y": [[8,2,14]],
    "names": ["Sleeping", "Awake"]
}
```


<br>

```sh
/bar?title=Bar+Chart
&subtitle=Sleeping+hours
&metric=hours
&theme=dark
&data={
    "x": ["Monday", "Friday", "Sunday"],
    "y": [[8,2,14]],
    "z": [[9,3,15]],
    "names": ["Sleeping", "Awake"]
}
```

## `GET /donut`


| `data`   | Required | Type     |
| :------- | :------- | :------- |
| `names`  | ◯        | []string |
| `values` | ◯        | []int    |

```sh
/donut?title=Donut+Chart
&data={
    "names": ["Monday", "Friday", "Saturday", "Sunday"],
    "values": [4, 6 ,7, 9]
}
```



## `GET /pie`

| `data`   | Required | Type     |
| :------- | :------- | :------- |
| `names`  | ◯        | []string |
| `values` | ◯        | []int    |


```sh
/pie?title=Pie+Chart
&subtitle=Sleeping+Hours
&data={
    "names": ["Monday", "Friday", "Saturday", "Sunday"],
    "values": [4, 6 ,7, 9]
}
```


## `GET /radar`

| `data`   | Required | Type     | Validation                         |
| :------- | :------- | :------- | :--------------------------------- |
| `names`  | ◯        | []string | `>=3`                              |
| `values` | ◯        | [][]int  | `count(names) == count(values[0])` |
| `labels` |          | []string |                                    |

```sh
/radar?title=Radar+Chart
&data={
    "names": ["Mon","Tue", "Wed", "Fri"],
    "labels": ["Work", "Relax", "Travel"],
    "values": [[1,2,3,4], [15,7,8,9], [15,17,5,7]]
}
```

## `GET /funnel`

| `data`   | Required | Type     |
| :------- | :------- | :------- |
| `names`  | ◯        | []string |
| `values` | ◯        | []int    |

```sh
/funnel?title=Radar+Chart
&data={
    "names": ["Mon","Tue", "Thu", "Fri", "Sat", "Sun"],
    "values": [100,80,60,40,20,10]
}
```


## `GET /table`

| `data`   | Required | Type       | Description  |
| :------- | :------- | :--------- | :----------- |
| `names`  | ◯        | []string   | `aka header` |
| `values` | ◯        | [][]string | `aka rows`   |

```sh
/table?title=Table+Chart
&data={
    "names": ["Branch","Code Coverage", "Quality"],
    "values": [["master","80","90"], ["develop","70","79"]]
}
```


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
