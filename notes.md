to fin out what is running into some port

### lsof -i:3000

to kill the stuff on that port

### kill $(lsof -t -i:3000)