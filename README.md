# repeat

**repeat** is a command line tool that outputs a user specified pattern a number of times, that is also specified by the user.

**repeat** is somewhat similar to the `seq` command found on Linux-based and Unix-like operating systems.
However, **repeat** types of patterns to repeat that are not possible with `seq`.


## Usage
Calling:
```
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
repeat --count=2 --plus=2 A b c d e
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
```


Calling:
```
# NOTE that --plus is set to NEGATIVE 2. (The previous example had positive 2.)
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


