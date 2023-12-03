`json2struct` is a Go package that exposes one function at the moment. This function is called `JSONToStruct` that takes a json as string and returns the equilvant go struct(s) that map to it. You can then use the structs as an argument to `Decode()` to deserialize your json.

Usage:

Download the package using the following command:

    go get github.com/walidsi/json2struct

Import it in your code:

    import "github.com/walidsi/json2struct"

Call:

    jsonString := "{\"id\": 2514, \"country\":\"EG\", \"sunrise\": 1701405191, \"sunset\": 1701442495}"
    jsonStructs, err := json2struct.JSONToStruct("root", jsonString)

I have created a simple web app that demonstrates the input and output. Check it at: http://json2struct.azurewebsites.net
