E-Biznes
---

### Zadanie 1 - [Docker](/zad1/) ([demo](/demos/zad1.mp4))
- [x] 3.0 - Obraz ubuntu z Pythonem w wersji 3.10 ([faf8c07](https://github.com/Tymec/e-biznes/commit/faf8c07576cba5b4c1ea737b856b8e4abe1598f5))
- [ ] 3.5 - Obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem
- [ ] 4.0 - Do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC SQLite w ramach projektu na Gradle
- [ ] 4.5 - Stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji przez CMD oraz gradle
- [ ] 5.0 - Dodać konfigurację docker-compose

### Zadanie 2 - [Scala](/zad2/) ([demo](/demos/zad2.mp4))
- [x] 3.0 - Należy stworzyć kontroler do Produktów ([6d4ed1e](https://github.com/Tymec/e-biznes/commit/6d4ed1efb9453faa1b9921fe1831265ab5c282b0))
- [ ] 3.5 - Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy
- [ ] 4.0 - Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy zgodnie z CRUD
- [ ] 4.5 - Należy aplikację uruchomić na dockerze oraz dodać skrypt uruchamiający aplikację via ngrok
- [ ] 5.0 - Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD

### Zadanie 3 - [Kotlin](/zad3/) ([demo](/demos/zad3.mp4))
- [x] 3.0 - Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord ([cf5babe](https://github.com/Tymec/e-biznes/commit/cf5babeac3e2664b4b6078f4b3f6b2ba4746c2ab))
- [ ] 3.5 - Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji
- [ ] 4.0 - Zwróci listę kategorii na określone żądanie użytkownika
- [ ] 4.5 - Zwróci listę produktów wg żądanej kategorii
- [ ] 5.0 - Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger, Webex

### Zadanie 4 - [Go](/zad4/) ([demo](/demos/zad4.mp4))
- [x] 3.0 - Należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD
- [x] 3.5 - Należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy)
- [ ] 4.0 - Należy dodać model Koszyka oraz dodać odpowiedni endpoint
- [ ] 4.5 - Należy stworzyć model kategorii i dodać relację między kategorią, a produktem
- [ ] 5.0 - Pogrupować zapytania w gorm'owe scope'y
