# Custom Handlers for Azure Functions

Write HTTP APIs in a given language, then you can now write functions in that language, for any of our triggers. We get requests for different languages through our various feedback channels, or in some cases users just want a little more control over the environment. There’s now a protocol for you to receive dispatch events using a lightweight HTTP server you provide (a “custom handler”), which can route things to individual functions.

## Getting Started

-	You just need a recent version of Functions and core tools, and you have everything you need, outside of any dependencies your target language requires.
-	The core of a handler is being able to write something that can receive a POST over http://127.0.0.1:port/functionName according to this request schema and response schema.
-	The FUNCTIONS_HTTPWORKER_PORT environment variable defines the port.
-	Beyond writing the handler against that schema, you’ll need to set the executable path and possibly a worker path (as in JS) in host.json.
    * See sample host.json that points to an executable [here](https://github.com/pragnagopa/functions-http-worker/blob/master/functions/host.json)
    * See sample host.json that points to a runtime and a worker path [here](https://github.com/pragnagopa/functions-http-worker/blob/master/functions/SampleHostJsonWithWorkerPath.json)
  
-	Some example handlers can be found in this repo
