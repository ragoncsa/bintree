# bintree

Simple coding challenge: render a binary tree on the terminal.

```bash
➜  bintree git:(main) go run main.go
INPUT: []int{74, 40, 75, 48, 86, 74, 92, 86, 44, 40, 75, 68, 97, 18, 13, 9, 29, 24, 75, 25}
74 -- 75 -- 86 -- 92 -- 97
 |     |     \ -- 86
 |     \ -- 75
 |           \ -- 75
 \ -- 40 -- 48 -- 74
       |     |     \ -- 68
       |     \ -- 44
       \ -- 40
             \ -- 18 -- 29
                   |     \ -- 24 -- 25
                   \ -- 13
                         \ -- 9
```