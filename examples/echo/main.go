package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	// "text/template"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
	owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("OWM_API_KEY")

type waHandler struct {
	wac       *whatsapp.Conn
	startTime uint64
}

func (wh *waHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "error caught in handler: %v\n", err)
}

// type printlocation struct {

// }

type User struct {
	Name string
	Bio  string
}

type CurrentTemperature struct {
	Temperature        float64
	WindSpeed          float64
	WindDirection      float64
	CloudAll           int
	Humidity           int
	Pressure           float64
	TempMin            float64
	TempMax            float64
	WeatherDescription string
	FeelsLike          float64
}

type currTempStr struct {
	Temperature        string
	WindSpeed          string
	WindDirection      string
	CloudAll           string
	Humidity           string
	Pressure           string
	TempMin            string
	TempMax            string
	WeatherDescription string
	FeelsLike          string
}

func getLocationData(message string) string {
	w, err := owm.NewCurrent("C", "en", apiKey) // celcius, english, openweathermap api key
	if err != nil {
		fmt.Println("Error while getting retrieving get weather data:")
		fmt.Fprintf(os.Stderr, "error getting location data: %v\n", err)
		log.Fatalln(err)
	}

	// wpol, errPol := owm.NewPollution(apiKey)
	// if errPol != nil {
	// 	fmt.Println("Error while getting retrieving get weather data:")
	// 	fmt.Fprintf(os.Stderr, "error getting location data: %v\n", errPol)
	// 	log.Fatalln(errPol)
	// }

	// tempcurrent := w.GetCurrent().Main.Temp + "°C"

	// var stringsplice  : = message.Text[1:]
	location := strings.ToLower(message)
	location = strings.TrimPrefix(location, "@w ")
	location = strings.TrimSpace(location)
	// fmt.Printf("%v\n", location)
	// fmt.Printf("%v\n", len(message))
	w.CurrentByName(location)
	fmt.Printf("%v\n", w.Weather)
	// w.

	// fmt.Printf(" Temperature\t - %t\n Wind Speed\t - %t\n Wind Direction\t - %t\n Cloud All\t - %t\n Humidity\t - %t\n Pressure\t - %t\n Temp Minumum\t - %t\n Temp Maximum\t - %t\n Weather Description\t - %t\n Feels Like\t - %t\n", w.Main.Temp, w.Wind.Speed, w.Wind.Deg, w.Clouds.All, w.Main.Humidity, w.Main.Pressure, w.Main.TempMin, w.Main.TempMax, w.Weather[0].Description, w.Main.FeelsLike)
	weatherDescription := ""
	if reflect.TypeOf(w.Weather) != nil && len(w.Weather) > 0 && reflect.TypeOf(w.Weather[0].Description) != nil {
		weatherDescription = w.Weather[0].Description
	} else {
		weatherDescription = "No Weather Data"
	}

	// if len(w.Weather) > 0 {
	// 	weatherDescription = "No weather data available"
	// } else {
	// 	weatherDescription = w.Weather[0].Description
	// }

	ct := CurrentTemperature{
		Temperature:        w.Main.Temp,
		WindSpeed:          w.Wind.Speed,
		WindDirection:      w.Wind.Deg,
		CloudAll:           w.Clouds.All,
		Humidity:           w.Main.Humidity,
		Pressure:           w.Main.Pressure,
		TempMin:            w.Main.TempMin,
		TempMax:            w.Main.TempMax,
		WeatherDescription: weatherDescription,
		FeelsLike:          w.Main.FeelsLike,
	}

	// ctstr := currTempStr{
	// 	Temperature: fmt.Sprintf("Temperature\t %v \n", w.Main.Temp),
	// 	WindSpeed: fmt.Sprintf("Wind Speed\t %v \n", w.Wind.Speed),
	// 	WindDirection: fmt.Sprintf("Wind Direction\t %v \n", w.Wind.Deg),
	// 	CloudAll: fmt.Sprintf("Cloud All\t %v \n", w.Clouds.All),
	// 	Humidity: fmt.Sprintf("Humidity\t %v \n", w.Main.Humidity),
	// 	Pressure: fmt.Sprintf("Pressure\t %v \n", w.Main.Pressure),
	// 	TempMin: fmt.Sprintf("Temp Minimum\t %v \n", w.Main.TempMin),
	// 	TempMax: fmt.Sprintf("Temp Maximum\t %v \n", w.Main.TempMax),
	// 	WeatherDescription: fmt.Sprintf("Weather Description \t %v \n", w.Weather[0].Description),
	// 	FeelsLike: fmt.Sprintf("Feels Like \t %v \n", w.Main.FeelsLike),
	// }

	// try {
	// tempcurrent := fmt.Sprintf(" Temperature\t - %v\n Weather Description\t - %v\n Feels Like\t - %v\n", w.Main.Temp, w.Weather[0].Description, w.Main.FeelsLike);
	// }	catch (error) {
	// 	return "Error"
	// }
	// tempcurrentStr := fmt.Sprintf(" Temperature\t - %v\n Wind Speed\t - %v\n Wind Direction\t - %v\n Cloud All\t - %v\n Humidity\t - %v\n Pressure\t - %v\n Temp Minumum\t - %v\n Temp Maximum\t - %v\n Weather Description\t - %v\n Feels Like\t - %v\n", w.Main.Temp, w.Wind.Speed, w.Wind.Deg, w.Clouds.All, w.Main.Humidity, w.Main.Pressure, w.Main.TempMin, w.Main.TempMax, w.Weather[0].Description, w.Main.FeelsLike)
	tempcurrentStr := fmt.Sprintf("Weather in %v\n\n Temperature\t  - %v %v\n Wind Speed\t - %v\n Wind Direction\t - %v\n Cloud All\t - %v\n Humidity\t - %v\n Pressure\t - %v\n Temp Minumum\t - %v\n Temp Maximum\t - %v\n Weather Description\t - %v\n Feels Like\t - %v\n Timezone - %v\n Location-Longitude - %v\n Localtion-Latitude %v\n", location, w.Main.Temp, w.Unit, w.Wind.Speed, w.Wind.Deg, w.Clouds.All, w.Main.Humidity, w.Main.Pressure, w.Main.TempMin, w.Main.TempMax, weatherDescription, w.Main.FeelsLike, w.Timezone, w.GeoPos.Longitude, w.GeoPos.Latitude)

	// u := User{"John", "a regular user"}
	// fmt.Println("--- csv 2 String ---")
	// fmt.Println(ct.bufStrTemperature(tempcurrentStr))
	// fmt.Println("--- End of csv 2 string ---")

	// ut, err := template.New("currenttemp").Parse("The user is {{ .Temperature }} \n and he is {{ .WindSpeed }}.")

	// if err != nil {
	// 	panic(err)
	// }

	// err = ut.Execute(os.Stdout, ct)

	// if err != nil {
	// 	panic(err)
	// }

	// TOdo write in string buffer and return the data in string format
	// return fmt.Sprintf("ut: %v\n", w)
	return ct.bufStrTemperature(tempcurrentStr)

}

func (ct CurrentTemperature) bufStrTemperature(fieldstobot string) string {
	buf := new(bytes.Buffer)
	buf.WriteString(fieldstobot)
	buf.WriteString("\n")

	return buf.String()
}

func (cstr currTempStr) check(field string) string {
	switch strings.ToLower(field) {
	case "Temperature":
		return cstr.Temperature
	case "Wind Speed":
		return cstr.WindSpeed
	case "Wind Direction":
		return cstr.WindDirection
	case "Cloud All":
		return cstr.CloudAll
	case "Humidity":
		return cstr.Humidity
	case "Pressure":
		return cstr.Pressure
	case "Temp Minimum":
		return cstr.TempMin
	case "Temp Maximum":
		return cstr.TempMax
	case "Weather Description":
		return cstr.WeatherDescription
	case "Feels Like":
		return cstr.FeelsLike
	default:
		return ""
	}
}

func (c CurrentTemperature) csv2String() string {
	// fmt.Println("--- csv %v", ctstr)
	buf := new(bytes.Buffer)
	v := reflect.ValueOf(c)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.String {
			// buf.WriteString()
			fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
			buf.WriteString("abcd" + f.String())
			buf.WriteString("\n")
		} else if f.Kind() == reflect.Int {
			fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
			buf.WriteString(fmt.Sprintf("%v", f.Int()))
			buf.WriteString("\n")
		} else if f.Kind() == reflect.Float64 {
			fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
			buf.WriteString(fmt.Sprintf("%v", f.Float()))
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

// func getWeatherData(message string) string {
// 	w, err := owm.NewForecast("C", "en", apiKey) // celcius, english, openweathermap api key
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	location := strings.ToLower(message)
// 	location = strings.TrimPrefix(location, "@w ")
// 	location = strings.TrimSpace(location)
// 	fmt.Printf("%v\n", location)
// 	fmt.Printf("%v\n", len(message))
// 	w.CurrentByName(location)
// 	fmt.Printf(" Temperature\t - %t\n Wind Speed\t - %t\n Wind Direction\t - %t\n Cloud All\t - %t\n Humidity\t - %t\n Pressure\t - %t\n Temp Minimum\t - %t\n Temp Maximum\t - %t\n Weather Description\t - %t\n Feels Like\t - %t\n", w.Main.Temp, w.Wind.Speed, w.Wind.Deg, w.Clouds.All, w.Main.Humidity, w.Main.Pressure, w.Main.TempMin, w.Main.TempMax, w.Weather[0].Description, w.Main.FeelsLike)
// 	return fmt.Sprintf("Temperature\t - %v\n Wind Speed\t - %v\n Wind Direction\t - %v\n Cloud All\t - %v\n Humidity\t - %v\n Pressure\t - %v\n Temp Minimum\t - %v\n Temp Maximum\t - %v\n Weather Description\t - %v\n Feels Like\t - %v\n", w.Main.Temp, w.Wind.Speed, w.Wind.Deg, w.Clouds.All, w.Main.Humidity, w.Main.Pressure, w.Main.TempMin, w.Main.TempMax, w.Weather[0].Description, w.Main.FeelsLike)
// }

// HandleTextMessage receives whatsapp text messages and checks if the message was send by the current
// user, if it does not contain the keyword '@echo' or if it is from before the program start and then returns.
// Otherwise the message is echoed back to the original author.
func (wh *waHandler) HandleTextMessage(message whatsapp.TextMessage) {
	fmt.Printf("%v\t%v\t\n", message.Info.FromMe, message.Text)
	if !strings.Contains(strings.ToLower(message.Text), "@w") || message.Info.Timestamp < wh.startTime {
		return
	}

	// if _, err := wh.wac.RevokeMessage(message.ID); err != nil {
	// 	fmt.Println(err)
	// }

	tempcurrent := getLocationData(message.Text)

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: message.Info.RemoteJid,
		},
		Text: tempcurrent,
	}

	if _, err := wh.wac.Send(msg); err != nil {
		fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
	}
	// Revoke the message to prevent it from being sent again
	// if _, errRevoke := wh.wac.RevokeMessage(message.Info.RemoteJid, message.Info.Id, message.Info.FromMe); errRevoke != nil {
	// 	fmt.Println(errRevoke)
	// }

	count := 4

	if _, loadMsgs := wh.wac.LoadMessages(message.Info.RemoteJid, message.Info.Id, count); loadMsgs != nil {
		fmt.Println("Loading messages: messages loaded.")
		fmt.Println(loadMsgs)
	}

	fmt.Printf("echoed message '%v' to user %v\n", message.Text, message.Info.RemoteJid)
}

func (wh *waHandler) HandleBatteryMessage(message whatsapp.BatteryMessage) {
	fmt.Printf("Plugged - %v\t Powersave- %v\t Percentage - %v\t\n", message.Plugged, message.Powersave, message.Percentage)
}

func (wh *waHandler) HandleLiveLocationMessage(message whatsapp.LiveLocationMessage) {
	fmt.Printf("Latitude - %v\t Longitude - %v\t\n AccuracyInMeters %v\t\n Degrees Clokwise From MagneticNorth - %v\t\n", message.DegreesLatitude, message.DegreesLongitude, message.AccuracyInMeters, message.DegreesClockwiseFromMagneticNorth)
}

func (wh *waHandler) HandleLocationMessage(message whatsapp.LocationMessage) {
	fmt.Printf("Latitude - %v\t Longitude - %v\t\n Name - %v\t Address -%v\t\n Url - %v\n", message.DegreesLatitude, message.DegreesLongitude, message.Name, message.Address, message.Url)
}

func downloadDocument(message whatsapp.DocumentMessage, wh *waHandler) {
	fmt.Printf("downloadDocument start kickstarted")
	data, err := message.Download()

	if err != nil {
		// if err != whatsapp {
		// 	log.Fatalf("error downloading media: %v\n", err)
		// 	return
		// }
		// if err != whatsapp.ErrMediaDownloadFailedWith410 && err != whatsapp.ErrMediaDownloadFailedWith404 {
		// 	fmt.Printf("%v\n", err)
		// 	return
		// }
		if _, err = wh.wac.LoadMediaInfo(message.Info.RemoteJid, message.Info.Id, strconv.FormatBool(message.Info.FromMe)); err == nil {
			data, err = message.Download()
			if err != nil {
				return
			}
		}
	}
	fmt.Printf("downloadDocument no error")

	filename := fmt.Sprintf("%v/%v.%v", os.TempDir(), message.Info.Id, strings.Split(message.Type, "/")[1])
	file, err := os.Create(filename)
	fmt.Print("downloadDocument created file")
	if err != nil {
		defer file.Close()
	}
	if err != nil {
		return
	}
	_, err = file.Write(data)
	if err != nil {
		return
	}
	fmt.Printf("downloadDocument written to file")
	log.Printf("%v %v\n\timage received, saved at:%v\n", message.Info.Timestamp, message.Info.RemoteJid, filename)
}

func (wh *waHandler) HandleDocumentMessage(message whatsapp.DocumentMessage) {
	fmt.Printf("%v\t%v\t\n", message.Info.FromMe, message.FileName)
	// if !message.Info.FromMe || message.Info.Timestamp < wh.startTime {
	// if !strings.Contains(strings.ToLower(message.FileName), "@d") || message.Info.Timestamp < wh.startTime {
	if message.Info.Timestamp < wh.startTime {
		return
	}

	fmt.Printf("HandleDocumentMessage Handler started")

	downloadDocument(message, wh)

	//f _, loadMsgs := wh.wac.

	fmt.Printf("Handle")
	if message.Info.FromMe {
		// Download document file

		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: message.Info.RemoteJid,
			},
			Text: "Attachment received",
		}
		if _, err := wh.wac.Send(msg); err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
		}
		fmt.Printf("echoed message '%v' to user %v\n", message.FileName, message.Info.RemoteJid)
	}

}

// func (wh *waHandler) HandlePdfMessage(message whatsapp.PdfMessage) {
// 	fmt.Printf("%v\t%v\t\n", message.Info.FromMe, message.FileName)
// }

func (wh *waHandler) HandleImageMessage(messageINGS whatsapp.ImageMessage) {
	fmt.Printf("%v\t%v\t\n", messageINGS.Info.FromMe, messageINGS.Caption)
	// if !messageINGS.Info.FromMe || messageINGS.Info.Timestamp < wh.startTime {
	if !strings.Contains(strings.ToLower(messageINGS.Caption), "@i") || messageINGS.Info.Timestamp < wh.startTime {
		return
	}

	if messageINGS.Info.FromMe {
		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: messageINGS.Info.RemoteJid,
			},
			Text: "Attachment received",
		}
		if _, err := wh.wac.Send(msg); err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
		}
		fmt.Printf("echoed message '%v' to user %v\n", messageINGS.Info, messageINGS.Info.RemoteJid)
	}

}

// func (wh *waHandler) HandleDeleteMessage(message whatsapp.DeleteMessage) {
// 	fmt.Printf("%v\t%v\t\n", message.Info.FromMe, message.MessageId)
// 	// if !message.Info.FromMe || message.Info.Timestamp < wh.startTime {

// }

func login(wac *whatsapp.Conn) error {
	session, err := readSession()
	if err == nil {
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return fmt.Errorf("restoring session failed: %v", err)
		}
	} else {
		qr := make(chan string)

		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()

		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("error during login: %v", err)
		}
	}

	if err = writeSession(session); err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}

	return nil
}

func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}

	UserHomeDir, err := os.UserHomeDir()
	fmt.Printf("user home directory is %v\n", UserHomeDir + os.TempDir())

	// file, err := os.Create(UserHomeDir + os.TempDir() + "/whatsappSession.gob")

	file, err := os.Open(UserHomeDir + os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err = decoder.Decode(&session); err != nil {
		return session, err
	}

	return session, nil
}

func writeSession(session whatsapp.Session) error {
	UserHomeDir, err := os.UserHomeDir()
		file, err := os.Create(UserHomeDir + os.TempDir() + "/whatsappSession.gob")
	// file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err = encoder.Encode(session); err != nil {
		return err
	}

	return nil
}

func main() {
	wac, err := whatsapp.NewConn(5 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return
	}

	wac.AddHandler(&waHandler{wac, uint64(time.Now().Unix())})

	if err = login(wac); err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return
	}

	<-time.After(60 * time.Minute)
}
