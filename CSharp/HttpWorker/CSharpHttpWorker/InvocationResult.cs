using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CSharpHttpWorker
{
    public class InvocationResult
    {
        public object ReturnValue { get; set; }

        public IDictionary<string, object> Outputs { get; set; } = new Dictionary<string, object>();

        public List<string> Logs { get; set; } = new List<string>();
    }
}
