# build
```
GOOS=linux GOARCH=386 go build ./server.go
```

# Flags

## 1. Dynamic flag
`curl IP:80/flag` > `haxagon{DYNAMIC_HASH}`

## 2. Static flag 
`curl IP:12345` > `flag1234`

## 3. Select flag
...

## 4. Exec check flag
`curl IP:12345/exec` should trigger the check

:)

<details>
  <summary>Spoiler warning</summary>
  
  Spoiler text. Note that it's important to have a space after the summary tag. You should be able to write any markdown you want inside the `<details>` tag... just make sure you close `<details>` afterward.
  
  ```javascript
  console.log("I'm a code block!");
  ```
  
</details>
