# repeat

**repeat** is a command line tool that outputs a user specified pattern a number of times, that is also specified by the user.

**repeat** is somewhat similar to the `seq` command found on Linux-based and Unix-like operating systems.
However, **repeat** types of patterns to repeat that are not possible with `seq`.


## Usage
Calling:
```
# The pattern to repeat is: A b c d e.
repeat A b c d e
```
Outputs:
```
A
b
c
d
e
```

Calling:
```
# By default the pattern will only be outputted once.
# Using --count=2 we make the pattern repeat two (2) times.
repeat --count=2 A b c d e
```
Outputs:
```
A
b
c
d
e
A
b
c
d
e
```


Calling:
```
# We can make a pattern partially repeat with the --plus switch.
# Using --plus=1 we make only first one (1) items from the pattern
# repeat AFTER all the full pattern repeats.
repeat --count=2 --plus=1 A b c d e
```
Outputs:
```
A
b
c
d
e
A
b
c
d
e
A
```


Calling:
```
# NOTE that --plus is set to NEGATIVE 2. (The previous example had positive 1.)
repeat --count=2 --plus=-2 A b c d e
```
Outputs:
```
A
b
c
d
e
A
b
c
d
e
A
b
c
```


