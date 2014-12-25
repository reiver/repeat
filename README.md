# repeat

**repeat** is a command line tool that outputs a user specified pattern a number of times, that is also specified by the user.

**repeat** is somewhat similar to the `seq` command found on Linux-based and Unix-like operating systems.
However, the type of patterns **repeat** can create are not possible with `seq`.


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
# Using --plus=1 we make only first one (1) item from the pattern
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
# NOTE that --plus is set to NEGATIVE 2 this time.
# (The previous example had positive 1.)
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


Calling:
```
# We use --separator=", " to stop each item from being put on its own line,
# and instead create a comma separated list.
repeat --count=2 --separator=", " A b c d e
```
Outputs:
```
A, b, c, d, e, A, b, c, d, e
```


Calling:
```
# We use a different separator here.
repeat --count=2 --separator=" + " A b c d e
```
Outputs:
```
A + b + c + d + e + A + b + c + d + e
```


Calling:
```
# We use the --format switch to transform the item from pattern before outputting it.
repeat --count=2 --format="f(%s)" A b c d e
```
Outputs:
```
f(A)
f(b)
f(c)
f(d)
f(e)
f(A)
f(b)
f(c)
f(d)
f(e)
```


Calling:
```
# We use both the --format switch and the separator switch.
repeat --count=2 --format="f(%s)" --separator=" ++ " A b c d e
```
Outputs:
```
f(A) ++ f(b) ++ f(c) ++ f(d) ++ f(e) ++ f(A) ++ f(b) ++ f(c) ++ f(d) ++ f(e)
```
