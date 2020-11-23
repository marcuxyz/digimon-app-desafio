import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.net.http.HttpResponse.BodyHandlers;
import java.nio.file.Files;
import java.nio.file.Path;
import java.time.Duration;
import java.time.Instant;
import java.util.List;

import com.google.gson.Gson;

class Main {

    class Digimon {
        private String name;
        private String img;
        private String level;

        public String getName() {
            return name;
        }

        public String getLevel() {
            return level;
        }

        public void setLevel(String level) {
            this.level = level;
        }

        public String getImg() {
            return img;
        }

        public void setImg(String img) {
            this.img = img;
        }

        public void setName(String name) {
            this.name = name;
        }
    }

    private static List<String> getNames() {
        try {
            List<String> lines = Files.readAllLines(Path.of("digimon.txt"));
            return lines;
        } catch (IOException e) {
            System.err.println("Erro ao ler o arquivo \"digimon.txt\"");
            System.exit(1);
            return null;
        }
    }

    private static Gson gson = new Gson();
    private static HttpClient client = HttpClient.newHttpClient();

    private static void request(String url) {
        HttpRequest request = HttpRequest.newBuilder().uri(URI.create(url)).build();

        try {
            HttpResponse<String> response = client.send(request, BodyHandlers.ofString());
            gson.fromJson(response.body(), Digimon[].class);
        } catch (IOException | InterruptedException e) {
            System.err.println("Erro ao realizar a requisição");
            System.exit(1);
        }
    }

    public static void main(String[] args) {
        Instant start = Instant.now();
        for (String name : Main.getNames()) {
            Main.request("https://digimon-api.vercel.app/api/digimon/name/" + name);
        }
        double duration = Duration.between(start, Instant.now()).toMillis() / 1000.;
        System.out.println(String.format("Duration (sequential): %f", duration));

        start = Instant.now();
        Main.getNames().parallelStream()
                .forEach(name -> Main.request("https://digimon-api.vercel.app/api/digimon/name/" + name));
        duration = Duration.between(start, Instant.now()).toMillis() / 1000.;
        System.out.println(String.format("Duration (parallel): %f", duration));
    }
}