<p align="center">
  <a href="https://github.com/kevincobain2000/instachart">
    <img alt="c2p" src="https://imgur.com/1RmqEWD.png" width="360">
  </a>
</p>
<p align="center">
  Chart to Image API interface for go-charts.
  <br>
  Data Driven READMEs.
  <br>

  <br>

</p>

**Quick Setup:** Easy setup. Add charts to your README.

**Hassle Free:** Create charts as images with single URL.

**Supports:** Choose from Line, bar, donut, pie, scatter, bubble, stacked bar, etc.

# Build Status


## Development

```sh
air
```

## API Docs

### Line Chart

Single date series

```sh
http://localhost:3000/line?data={
    "x": [["2022-12-23","2022-12-24","2023-12-25"]],
    "y": [[1,2,3]]
}
```

Multiple date series

```sh
http://localhost:3000/line?data={
    "x": [["2022-12-23","2022-12-24","2023-12-25"], ["2022-12-23","2022-12-28","2023-12-30"]],
    "y": [[1,2,3], [2,5,3]]
}
```

Continuous Series

```sh
http://localhost:3000/line?data={
    "x": [["10","20","30"]],
    "y": [[1,2,3]]
}
```

Multiple Continuous Series

```sh
http://localhost:3000/line?data={
    "x": [["10","20","30"], ["10","20","30"]],
    "y": [[1,2,3], [10,20,30]]
}
```

### Bar Chart

```sh
http://localhost:3000/bar?title=sleeping+hours
&y_label=hours
&data={ "x": ["Monday","Sunday"], "y": [8, 14] }
```