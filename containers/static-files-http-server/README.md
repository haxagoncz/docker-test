## Co je pod povrchem?

#### About
```
Najdi na stránce vlajku, která není moc dobře ukryta.
```

#### Hint
```
Zkontroluj konzoli na webu.
```

#### Result
```
ssps{CONSOLE_LOG}
```

#### Install
```
docker build -t console_log .
docker run -d --name console_log -p 7575:80 console_log
```

- port:7575
