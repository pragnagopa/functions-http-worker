using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace CSharpHttpWorker.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class HttpTriggerController : ControllerBase
    {
        private ILogger<ValuesController> _logger = null;
        public HttpTriggerController(ILogger<ValuesController> logger)
        {
            _logger = logger;
        }

        // GET api/values/5
        [HttpGet]
        public ActionResult<string> Get()
        {
            return "hello from c# worker";
        }

        // POST api/values
        [HttpPost]
        public ActionResult<string> Post(string value)
        {
            var invocationRequest = this.HttpContext;
            _logger.LogInformation($"queryparam:{invocationRequest.Request.Query["name"]}");
            return "HelloWorld from c# worker";
        }
    }
}