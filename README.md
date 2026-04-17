# bird

A command-line tool for tired brains.

```
$ bird
You logged in. That counts.

$ bird
    __
  (o >
   \_\
   "" ""

$ bird
Pigeons can do basic math.
```

No flags. No args. No config. No network.

You type `bird`. Something appears. Different each time.

## Install

```
go install github.com/DavidLiedle/bird@latest
```

Or clone and `go build`.

## Why

One day in 2016 I got to the office so tired I just typed `bird` into
my terminal. My brain had offered me a word, and the shell offered me
nothing back. This is the thing that should have been there.

Your tired brain can produce exactly four letters. Those four letters
should give something back.

## What it does

Three behaviors, chosen at random:

- A small ASCII bird.
- A calm sentence, written for the person at the keyboard.
- One fact about a real bird.

That's the whole product. It ignores any arguments you pass it, because
the tired brain is allowed to type nonsense.

## License

MIT.
