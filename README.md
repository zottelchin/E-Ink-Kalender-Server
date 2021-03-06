# E-Ink-Kalender-Server
Ich wollte einen E-Ink Kalender mit einem ESP-8266 bauen. Da es noch keinen Code gab, der genau das machte, was ich wollte, habe ich angefangen nach ähnlichen Projekten zu suchen.
[Link zum Client Code (ESP8266)](https://github.com/zottelchin/E-Ink-Kalender-Client)

**Wichtig: Der Server darf nicht hinter einem https Proxy laufen, da sonst der ESP nicht drauf zugreifen kann.**

# Hardware:
- ESP8266 Waveshare ESP Driver Board [Waveshare-Seite](https://www.waveshare.com/wiki/E-Paper_ESP8266_Driver_Board)
- Waveshare 4,2" ESP Raw Panel [Waveshare-Seite](https://www.waveshare.com/wiki/4.2inch_e-Paper_Module)
## Case
Den Display möchte ich gerne in einem Holzgehäuse auf meinen Schreibtisch stellen. Die Idee besteht schon, ist aber noch nicht umgesetzt. Sobald das der Fall ist, wird das hier nachgetragen.

## 1. Versuch
Als ersten Versuch habe ich den [Code](https://github.com/doctormord/ESP8266_EPD_Weather_Google_Calendar) von @doctormord angeguckt und einfach den Wetterpart gelöscht. Leider hat das bei mir nicht funktioniert, da der Heap meines ESPs zu klein war, was ihn zum Neustart gezwungen hat.

## 2. Versuch
Da der erste Versuch ja gescheitert ist, kam dann die Überlegung, die Daten selber dem ESP zuzuführen, dadurch benötigt man nicht den Google Redirect Code. So entstand der Code in diesem Repo. Da ich in GO schon öfters gearbeitet habe, habe ich kurzerhand einen Server geschrieben, der mir meine Kalederdaten als String liefert.
Ein schöner Nebeneffekt des Selberschreibens ist, dass ich direkt .ics Dateien verarbeiten kann und nicht auf Google Kalender limitiert bin. 
### Server Funktion
Bekommt der Server einen GET Request (URL kann frei gewählt werden, wird auch noch in eine Configuration ausgelagert werden), ruft er die ics Dateien ab und filtert alle vergangenen Ereignisse raus und sortiert den Rest nach dem Start-Zeitpunkt. Aus dieser Liste werden die 5 aktuellsten Einträge in einen String überfühert. Diesem wird noch das aktuelle Datum angefügt und dann als Antwort gesendet.
#### verwendete Bibliotheken anderer Nutzer:
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [ics-golang Fork](https://github.com/rjhorniii/ics-golang/tree/ical2org)
### Client Funktion
Der Client verbindet sich mit dem WLAN und ruft dann die voreingestellte Seite auf und bekommt so die Daten. Der String wird aufgesplitte (an den ';') und dann Zeilenweise angezeigt.
#### Bilder 
##### Startbildschirm:
![Startbildschirm](/Bilder/IMG_20181118_220549.jpg)
##### Wenn der Server nicht erreicht werden kann, dann wird dieser Bildschrim angezeigt:
![Keine Serververbindung](/Bilder/IMG_20181118_221312.jpg)
##### Übersicht über die 5 aktuellsten Termine
![Kalenderansicht](/Bilder/IMG_20181118_220928.jpg)
#### verwendete Bibliotheken anderer Nutzer:
- [EPD Driver](https://github.com/ZinggJM/GxEPD)
- Adafruit Fonts
## Pläne für die Zukunft
### Server
- [x] Key an Route hängen, damit Daten halbwegs sicher sind
- [x] .ics Links aus Datei lesen
### Client
- [x] UTF-8 Unterstützung U8g2-Font ist aktuell der Plan 
