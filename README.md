# logger2
golang logger for rdpv2. leverages Uber zap

### usage
```
func doSomeStuff() {
    // these lines only show when Client is initialized with NewLocal()
	logger2.Client.Debug("the debug stuff is hererererererererre", zap.String("d", "debugggggg"))
	
	// adds an info-level entry to stacktrace array
	logger2.Info(
		map[string]interface{}{
			"ValA": "A",
		},
		logger2.WhereAmI())
	//adds an error-level entry to stacktrace
	logger2.Error(
		map[string]interface{}{
			"ValA": "A",
		},
		logger2.WhereAmI())
}
```
### For Production
```
func main() {
    // initializes a stacktrace array
	logger2.InitStackTrace()

	logger2.Client = logger2.NewProduction().With(
		zap.Any("product", "testP"),
		zap.String("createAt", time.Now().UTC().Format(time.RFC1123)),
	)
	defer func() {
		logger2.Client.Sync()
	}()
}

```
### For Local Development
```
func main() {
    // initializes a stacktrace array
	logger2.InitStackTrace()

	logger2.Client = logger2.NewProduction().With(
		zap.Any("product", "testP"),
		zap.String("createAt", time.Now().UTC().Format(time.RFC1123)),
	)
	defer func() {
		logger2.Client.Sync()
	}()

	doSomeStuff()
	// continue with calling of functions
	
	//at the end, call logger2.Client.Info() to send the collected stack trace to cwl
	logger2.Client.Info(
	    "summary of TestLoggerDebugLevel", 
	    zap.Any("stackTraceArray", logger2.STA),
	    )
	
}

```
