# Timer
Simple time management tool for the terminal.

## Setup
Compile with `go build`.

### Bash

#### Add to .bashrc:
```bash
export TIMER_PATH="$HOME/path/to/executable"
if [ -d $TIMER_PATH= ]; then
    export PATH="$PATH:$TIMER_PATH"
fi
```

## Usage
One flag to rule them all.

Execute with `timer [flags]`.

Flags:
- `-time`: Time intervals as a comma-separated list.

### Controls

- The timer has to be started manually.
- Press `space` to start or stop the timer.
- Interrupt with `ctrl + c` to close the program.

### Examples

#### Plain 2 min 30 sec timer:
    ```sh
    timer -time 2m30s
    ```

#### Two 10 min sessions with a 5 min break:
    ```sh
    timer -time 10m,5m,10m
    ```
    or
    ```sh
    timer -time "10m, 5m, 10m"
    ```
    or might as well
    ```sh
    timer -time "10m , 5m" -time 10m
    ```

#### Easy pomodoro timer:
    ```sh
    timer -time 25m,5m,25m,5m,25m,5m,25m
    ```
    or use it as an alias
    ```bash
    alias pomodoro="timer -time 25m,5m,25m,5m,25m,5m,25m"
    ```
