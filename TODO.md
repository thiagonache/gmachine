[-] How to indicate some error in a G-machine program? Error register?
[X] Rune literals: SETA 'A'
[-] Monitor/debugger

```
P: 0 A: 0 I: 0 Op: SETA

> R
P: 3 A: 1 I: 0 Op: HALT

> ?
R: Run till next breakpoint
S: Set register
M: Set memory
Q: Quit

> M
Location? 3
Data? 1

P: 3 A: 1 I: 0 Op: NOOP

> R
Register? A
Data? 99

P: 3 A: 99 I: 0 Op: HALT
```
