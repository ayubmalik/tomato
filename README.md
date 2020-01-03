# 🍅 tomato
Tomato is a simple [pomodoro](https://en.wikipedia.org/wiki/Pomodoro_Technique) timer built using Go.

## Usage

After building, run the `tomato` command.

By default tomato will start a timer for 30 minutes. After the timer has expired the screen (terminal) will flash and exit.

To change the duration of the timer use the time flag with a [duration string](https://golang.org/pkg/time/#ParseDuration) e.g. 30m, 1h or 90s etc.

```
tomato -time 5m
```

# TODO
 * Add repeats and interval
