using System;
using System.IO;
using System.Net.Http;
using System.Threading.Tasks;
using System.Diagnostics;
using System.Linq;
using System.Collections.Generic;

class Program
{
    private static readonly HttpClient client = new HttpClient()
    {
        BaseAddress = new Uri("https://digimon-api.vercel.app/api/digimon/name/")
    };

    static async Task Main(string[] args)
    {
        var watch = Stopwatch.StartNew();
        var digimonTasks = new List<Task<string>>(209);

        const int BufferSize = 128;
        using (var fileStream = File.OpenRead("digimon.txt"))
        using (var streamReader = new StreamReader(fileStream, System.Text.Encoding.UTF8, true, BufferSize))
        {
            String digimon;
            while ((digimon = streamReader.ReadLine()) != null)
            {
                var task = GetData(digimon);
                digimonTasks.Add(task);
            }
        }

        var strings = await Task.WhenAll(digimonTasks).ConfigureAwait(false);
        watch.Stop();

        // Script executed in 1074 milliseconds.
        foreach(var teste in strings)
        {
            Console.WriteLine(teste);
        }
        
        Console.WriteLine($"Script executed in {watch.ElapsedMilliseconds} milliseconds.");
        Console.ReadKey();
    }

    static async Task<string> GetData(string digimon)
    {
        using (var res = await client.GetAsync(digimon).ConfigureAwait(false))
        {
            using (var content = res.Content)
            {
                return await content.ReadAsStringAsync().ConfigureAwait(false);
            }
        }
    }
}