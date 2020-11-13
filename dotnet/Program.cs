using System;
using System.IO;
using System.Net.Http;
using System.Threading.Tasks;
using System.Diagnostics;

class Program
{
    static void Main(string[] args)
    {
        var watch = Stopwatch.StartNew();

        foreach (string digimon in File.ReadLines("digimon.txt"))
        {
            Task.Run(() => GetData(digimon)).Wait();
        }

        watch.Stop();
        // Script executed in 729 milliseconds.
        Console.WriteLine($"Script executed in {watch.ElapsedMilliseconds} milliseconds.");
    }
    static async void GetData(string digimon)
    {
        HttpClient client = new HttpClient();

        string baseUrl = "https://digimon-api.vercel.app/api/digimon/name/";
        
        HttpResponseMessage res = await client.GetAsync(baseUrl + digimon);
        HttpContent content = res.Content;

        string data = await content.ReadAsStringAsync();

        Console.WriteLine(data);
    }
}