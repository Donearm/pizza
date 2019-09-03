package main

////////////////////////////////////////////////////////////////////////////////
//
// Copyright (c) 2019, Gianluca Fiore
//
////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"encoding/json"
	"os"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/mziia/geogoth"
)

const GEOJSON_FILE = "pizzerias.geojson"

type Pizzeria struct {
	address string
	city string
	country string
	name string
	website string
}

var  pizzaList = []Pizzeria {
	{"623 E. Adams Street", "Phoenix", "USA", "Pizzeria Bianco", "http://www.pizzeriabianco.com"},
	{"1504 Shattuck Avenue", "Berkeley", "USA", "The Cheese Board Collective", "http://cheeseboardcollective.com"},
	{"2995 Shattuck Avenue", "Berkeley", "USA", "Emilia's", "http://emiliaspizzeria.com"},
	{"101 Salem Street", "Chicago", "USA", "Celestino's", "http://www.celestinosnypizza.com"},
	{"641 North Highland Avenue", "Los Angeles", "USA", "Pizzeria Mozza", "http://www.pizzeriamozza.com"},
	{"6211 Shattuck Avenue", "Oakland", "USA", "Nicks", "http://oaklandstylepizza.com"},
	{"448 S. California Avenue", "Palo Alto", "USA", "Terun", "http://terunpizza.com"},
	{"1956 Bacon Street", "San Diego", "USA", "Pizza Port Ocean Beach", "http://www.pizzaport.com"},
	{"4046 30th Street", "San Diego", "USA", "Sicilian Thing Pizza", "http://www.sicilianthingpizza.com"},
	{"1570 Stockton Street", "San Francisco", "USA", "Tony's Pizza Napoletana", "http://tonyspizzanapoletana.com"},
	{"210 11th Street", "San Francisco", "USA", "Una Pizza Napoletana", "http://www.unapizza.com"},
	{"2723 Wilshire Boulevard", "Santa Monica", "USA", "Milo & Olive", "http://www.miloandolive.com"},
	{"1730 Pearl Street", "Boulder", "USA", "Pizzeria Locale", "http://localeboulder.com"},
	{"1610 16th Street", "Denver", "USA", "Lucky Pie Pizza & Taphouse", "http://www.lukypiepizza.com/lodo"},
	{"500 16th Street", "Denver", "USA", "Grimaldi's Pizzeria", "http://www.grimaldispizzeria.com"},
	{"2500 Larimer Street", "Denver", "USA", "Cart Driver", "http://cart-driver.com"},
	{"858 State Street", "New Haven", "USA", "Da Legna", "http://dalegna.com"},
	{"874 State Street", "New Haven", "USA", "Modern Apizza", "http://www.modernapizza.com"},
	{"237 Wooster Street", "New Haven", "USA", "Sally's Apizza", "http://www.sallysapizza.com"},
	{"3715 Macomb Street NW", "Washington", "USA", "2 Amys", "http://2amyspizza.com"},
	{"1400 Irving Street NW #103", "Washington", "USA", "Pete's New Haven Pizza", "http://petesapizza.com"},
	{"1245 W University Avenue", "Gainesville", "USA", "Leonardo's By The Slice", "http://www.leonardosgainesville.com"},
	{"1800 NE 23rd Avenue", "Gainesville", "USA", "Satchel's", "http://www.satchelspizza.com"},
	{"11551 University Boulevard", "Orlando", "USA", "Lazy Moon", "http://www.lazymoonpizza.com"},
	{"591 Edgewood Avenue", "Atlanta", "USA", "Ammazza", "http://ammazza.com"},
	{"1093 Hemphill Avenue", "Atlanta", "USA", "Antico Pizza Napoletana", "http://littleitalia.com"},
	{"1321 West Grand Avenue", "Chicago", "USA", "Coalfire", "http://coalfirechicago.com"},
	{"1769 West Sunnyside Avenue", "Chicago", "USA", "Spacca Napoli", "http://www.spaccanapolipizzeria.com"},
	{"1791 East 10th Street", "Bloomington", "USA", "Pizza X", "http://www.pizzaxbloomington.com"},
	{"30 S Meridian Street", "Indianapolis", "USA", "Napolese", "http://napolesepizzeria.com"},
	{"111 Chelsea Street", "Boston", "USA", "Santarpio's Pizza", "http://www.santarpiospizza.com"},
	{"119 N 4th Street", "Minneapolis", "USA", "Pizza Luce", "https://pizzaluce.com"},
	{"5557 South Xerxes Avenue", "Minneapolis", "USA", "Pizzeria Lola", "http://www.pizzerialola.com"},
	{"210 E Hennepin Avenue", "Minneapolis", "USA", "Punch Neapolitan Pizza", "http://www.punchpizza.com"},
	{"140 S Green Valley Parkway", "Henderson", "USA", "Settebello", "http://settebello.net"},
	{"911 Kingsley Street", "Asbury Park", "USA", "Porta", "http://pizzaporta.com"},
	{"147 Sloan Avenue", "Hamilton", "USA", "DeLorenzo's Pizza", "https://www.delorenzospizza.com/"},
	{"135 Newark Avenue", "Jersey City", "USA", "Porta", "http://pizzaporta.com"},
	{"275 Grove Street", "Jersey City", "USA", "Razza", "http://www.razzanj.com"},
	{"510 Central Avenue SE", "Albuquerque", "USA", "Farina", "http://www.farinapizzeria.com"},
	{"3410 Central Avenue SE", "Albuquerque", "USA", "Slice Parlor", "http://www.sliceparlor.com"},
	{"485 Lorimer Street", "New York", "USA", "Forcella", "https://www.forcellaeatery.com/"},
	{"575 Henry Street", "New York", "USA", "Lucali", "http://www.lucali.com"},
	{"60 Greenpoint Avenue", "New York", "USA", "Paulie Gee's", "http://pauliegee.com"},
	{"261 Moore Street", "New York", "USA", "Roberta's", "http://www.robertaspizza.com"},
	{"435 Halsey Street", "New York", "USA", "Saraghina", "http://www.saraghinabrooklyn.com"},
	{"278 Bleecker Street", "New York", "USA", "John's of Bleeker St", "http://www.johnsbrickovenpizza.com"},
	{"32 Spring Street", "New York", "USA", "Lombardi's", "http://www.firstpizza.com"},
	{"349 East 12th Street", "New York", "USA", "Motorino", "http://www.motorinopizza.com"},
	{"27 Prince Street", "New York", "USA", "Prince Street Pizza", "http://www.princestreetpizzanyc.com"},
	{"631 Haywood Road", "Asheville", "USA", "Standard Pizza Co.", "http://www.standardpizzacoasheville.com"},
	{"105 East Chapel Hill Street", "Durham", "USA", "Pizzeria Toro", "http://www.pizzeriatoro.com"},
	{"1842 Wake Forest Road", "Raleigh", "USA", "Capital Creations Gourmet Pizza", "http://capitalcreations.com"},
	{"240 North Liberty Street", "Powell", "USA", "Brooklyn Pizza", "http://www.brooklynpizzapowell.com"},
	{"4741 SE Hawthorne Boulevard", "Portland", "USA", "Apizza Scholls", "http://apizzascholls.com"},
	{"338 NW 21st Avenue", "Portland", "USA", "Ken's Artisan", "http://kensartisan.com"},
	{"701 South 50th Street", "Philadelphia", "USA", "Dock Street Brewing Co.", "http://www.dockstreetbeer.com"},
	{"4116 Ridge Avenue", "Philadelphia", "USA", "In Riva", "http://www.in-riva.com"},
	{"611 South 7th Street", "Philadelphia", "USA", "Nomad Pizza", "http://www.nomadpizzaco.com"},
	{" 2112 Murray Avenue", "Pittsburgh", "USA", "Aiello's Pizza", "http://aiellospizza.com"},
	{"116 West Market Street", "Pottsville", "USA", "Roma Pizza", "http://www.romapizzapottsville.com"},
	{"4437 Kingston Pike", "Knoxville", "USA", "Hard Knox Pizzeria Bearden", "https://hardknoxpizza.com"},
	{"115 16th Avenue South", "Nashville", "USA", "Desano Pizza Bakery", "http://desanopizza.com"},
	{"1502 South 1st Street", "Austin", "USA", "40 North", "http://www.40northpizza.com"},
	{"1305 West Oltorf", "Austin", "USA", "The Austin Beer Garden Brewing Company", "http://theabgb.com"},
	{"507 San Jacinto Street", "Austin", "USA", "The Backspace", "http://thebackspace-austin.com"},
	{"1519 East Cesar Chavez Street", "Austin", "USA", "Bufalina", "http://www.bufalinapizza.com"},
	{"417 Travis Street", "Houston", "USA", "Frank's Pizza", "http://frankspizza.com"},
	{"4920 S. Genesee Street", "Seattle", "USA", "Flying Squirrel", "http://www.flyingsquirrelpizza.com"},
	{"8310 5th Avenue NE", "Seattle", "USA", "Flying Squirrel", "http://www.flyingsquirrelpizza.com"},
	{"5701 Airport Way South", "Seattle", "USA", "Flying Squirrel", "http://www.flyingsquirrelpizza.com"},
	{"525 Rainier Avenue South", "Seattle", "USA", "Humble Pie", "http://humblepieseattle.com"},
	{"316 Virginia Street", "Seattle", "USA", "Serious Pie", "http://seriouspieseattle.com"},
	{"913 East Pike Street", "Seattle", "USA", "Via Tribunali", "http://viatribunali.com"},
	{"912 East Johnson Street", "Madison", "USA", "Salvatore's Tomato Pies", "http://salvatorestomatopies.com"},
	{"Marktplatz 2", "Lambach", "Austria", "La Stella del Sud", "https://www.instagram.com/la.stella.del.sud.lambach/"},
	{"Rruga  Brigada e VIII", "Tirana", "Albania", "Pepper", "http://pepper-tirana.com"},
	{"Place Saint-Josse 8", "Bruxelles", "Belgium", "La Piola Pizza", "http://www.lapiolapizza.com"},
	{"Rue Lebeau 75", "Bruxelles", "Belgium", "La Pizza è Bella", "https://lapizzaebella.be/en"},
	{"Place Saint-Josse 8", "Bruxelles", "Belgium", "La Piola Pizza", "http://www.lapiolapizza.com"},
	{"Rue Lebeau 75", "Bruxelles", "Belgium", "La Pizza è Bella", "https://lapizzaebella.be/en"},
	{"Pensstraat 6", "Leuven", "Belgium", "La Vecchia Napoli", "http://www.lavecchianapoli.be"},
	{"11 Clinton Street", "Toronto", "Canada", "Bitondo Pizzeria", ""},
	{"155 University Avenue", "Toronto", "Canada", "Pizzeria Libretto", "http://pizzerialibretto.com"},
	{"2336 Lake Shore Boulevard West", "Toronto", "Canada", "FBI Pizza", "http://www.fbipizza.com"},
	{"214 Church Street", "Moncton", "Canada", "Zio's", "https://m.facebook.com/pages/Zios-Pizza/444640382240517"},
	{"Revoluční 655/1", "Prague", "Czechia", "Pizza Nuova", "http://www.pizzanuova.cz"},
	{"Guldbergsgade 29", "Copenhagen", "Denmark", "Bæst", "http://baest.dk"},
	{"Gammel Strand 42", "Copenhagen", "Denmark", "Luca", "https://www.iloveluca.dk"},
	{"Høkerboderne 9-15", "Copenhagen", "Denmark", "Mother", "https://mother.dk"},
	{"Õle 33", "Tallin", "Estonia", "Kaja Pizza Köök", "http://kajapizza.ee"},
	{"Kalevankatu 6", "Helsinki", "Finland", "Putte's", "http://puttes.fi"},
	{"3 Quai Saint-Pierre", "Cannes", "France", "La Pizza Cresci", "http://maison-cresci.fr/en"},
	{"10 Rue Dancourt", "Paris", "France", "Bijou", "https://bijou-paris.fr"},
	{"45 Rue Brancion", "Paris", "France", "Guillame Grasso", "https://www.guillaume-grasso.com"},
	{"78 Rue de Charonne", "Paris", "France", "Louie Louie", "https://www.louielouie.paris"},
	{"108 Avenue de Villiers", "Paris", "France", "Manhattan Terrazza", "https://www.manhattanterrazza.fr"},
	{"107 Boulevard Richard Lenoir", "Paris", "France", "Ober Mamma", "https://www.bigmammagroup.com/en/accueil"},
	{"111 Rue Réaumur", "Paris", "France", "Pizzeria Popolare", "https://www.bigmammagroup.com"},
	{"8 Rue Rossini", "Paris", "France", "Rossini", "http://www.yelp.com/biz/rossini-paris-3"},
	{"88 Boulevard de Belleville", "Paris", "France", "Tripletta", "https://triplettabelleville.fr"},
	{"Lychener Straße 2", "Berlin", "Germany", "PizzaNostra", "https://www.pizzanostra.de"},
	{"Hobrechtstraße 57", "Berlin", "Germany", "Gazzo", "https://www.gazzopizza.com"},
	{"Hauptstraße 85", "Berlin", "Germany", "Malafemmena", "http://malafemmena.restaurant"},
	{"Straßmannstraße 21", "Berlin", "Germany", "Pomodorino", "https://www.pomodorino.de/pomodorino"},
	{"Templiner Straße 7", "Berlin", "Germany", "Standard Serious Pizza", "http://www.standard-berlin.de"},
	{"Sredzkistraße 49", "Berlin", "Germany", "Trattoria Toscana", "http://www.toscana-tempelhof.de"},
	{"Weserstraße 14", "Franfkurt am Main", "Germany", "Montana", "http://www.montana-pizzeria.de"},
	{"Occamstraße 11", "Munchen", "Germany", "60 Secondi", "https://www.60-seconds.de"},
	{"208 Rathmines Road Lower", "Dublin", "Ireland", "Manifesto", "https://www.manifestorestaurant.ie"},
	{"Vicolo San Giovanni Battista 3", "Caiazzo", "Italy", "Pepe In Grani", "http://www.pepeingrani.it"},
	{"Via Antonio Vivaldi 23", "Caserta", "Italy", "I Masanielli", "https://www.facebook.com/masaniellisasamartucci"},
	{"Piazza Matteotti 40", "Caserta", "Italy", "La Loggetta", "https://www.facebook.com/PIZZERIALALOGGETTALAB"},
	{"Via Cesare Battisti 46", "Caserta", "Italy", "Tre Farine", "http://www.trefarine.it"},
	{"Via Montevideo 4", "Milan", "Italy", "Olio a Crudo", "https://www.sorbillo.it/pizzeria-olio-a-crudo"},
	{"Piazza Sannazaro 201/c", "Naples", "Italy", "50 Kalò", "http://www.50kalo.it"},
	{"Via dei Tribunali 94", "Naples", "Italy", "Di Matteo", "http://www.pizzeriadimatteo.com"},
	{"Via dei Tribunali 32", "Naples", "Italy", "Gino e Toto Sorbillo", "http://www.sorbillo.it"},
	{"Via Cesare Sersale 1-3", "Naples", "Italy", "L'Antica Pizzeria da Michele", "http://www.damichele.net"},
	{"Via Michelangelo da Caravaggio 53", "Naples", "Italy", "La Notizia", "http://www.pizzarialanotizia.com"},
	{"Via Materdei 27/28", "Naples", "Italy", "Starita", "https://www.pizzeriastarita.it"},
	{"Via Nuova Agnano 1", "Pozzuoli", "Italy", "10 Diego Vitagliano", "http://www.10pizzeria.it"},
	{"Via Campi Flegrei 13", "Pozzuoli", "Italy", "Biga280", "http://www.biga280.it"},
	{"Via Appia Nuova 1095", "Rome", "Italy", "Angelo Pezzella", "http://www.angelopezzella.it"},
	{"Via Mantova 5", "Rome", "Italy", "Berberè", "https://www.berperepizza.it/en"},
	{"Via San Biagio Platani 320", "Rome", "Italy", "I Quintili", "https://www.facebook.com/Iquintili1"},
	{"Via Federico Ozanam 30-32", "Rome", "Italy", "La Gatta Mangiona", "http://www.lagattamangiona.com"},
	{"Via del Lavatore 91", "Rome", "Italy", "Piccolo Buco", "https://www.pizzeriapiccolobuco.it"},
	{"Via della  Meloria 43", "Rome", "Italy", "Pizzarium", "http://www.bonci.it"},
	{"Via Torrione 36", "Salerno", "Italy", "Capri", "http://www.capripizzeriasalerno.it"},
	{"Via Andrea Sabatini 4", "Salerno", "Italy", "Resilienza", "https://www.facebook.com/PizzeriaResilienza"},
	{"Via Camporosolo 11", "San Bonifacio", "Italy", "I Tigli", "http://www.pizzeriaitigli.it"},
	{"Largo Arso 10-16", "San Giorgio a Cremano", "Italy", "Pizzeria Fratelli Salvo", "http://www.salvopizzaioli.it"},
	{"Via Ponte 55A", "San Martino Buon Albergo", "Italy", "Saporè", "http://www.saporeverona.it"},
	{"NS Azabu Juban  Bldg. 1F, 3-6-2", "Tokyo", "Japan", "Pizza Strada", "http://www.pizzastrada.jp"},
	{"279 Saint Paul Street", "Valletta", "Malta", "La Vecchia Taranto", ""},
	{"Henrik Ibsens Gate 60a", "Oslo", "Norway", "Vinoteket", "http://vinoteket.no"},
	{"Grodzka 63", "Kraków", "Poland", "Fiorentina", "https://fiorentina.com.pl"},
	{"Krakowska 27", "Kraków", "Poland", "Nolio", "https://nolio.pl"},
	{"Kalwaryjska 32", "Kraków", "Poland", "Pizzeria 00", "http://pizzeria00.pl"},
	{"Kazimierza Brodzińskiego 4", "Kraków", "Poland", "Primo Italian", "http://www.primoitalian.pl"},
	{"Topiel 12", "Warsaw", "Poland", "Ave Pizza", "https://www.avepizza.pl"},
	{"Jarosława Dąbrowskiego 27", "Warsaw", "Poland", "Ciao a Tutti Due", "https://www.facebook.com/ciaotuttipizza/"},
	{"Chmielna 13a", "Warsaw", "Poland", "Mąka i Woda", "https://www.facebook.com/MakaiWoda"},
	{"Rua da Artilharia 1 16b", "Lisbon", "Portugal", "Forno d'Oro", "http://www.fornodoro.pt"},
	{"Strada Visarion 10", "Bucharest", "Romania", "Animaletto", "http://www.animaletto.ro"},
	{"Ulitsa Kiyevskaya 2", "Moscow", "Russia", "Eataly Moscow", "https://www.eataly.ru"},
	{"Smolenskaya Square 3", "Moscow", "Russia", "Luciano", "http://www.lucianomoscow.ru"},
	{"Dunajská 25", "Bratislava", "Slovakia", "Basilicò", "https://basilico-italian-restaurant.business.site"},
	{"Carrer de Balmes 193", "Barcelona", "Spain", "La Balmesina", "http://www.labalmesina.com"},
	{"Calle de Sta Engracia 48", "Madrid", "Spain", "Grosso Napoletano", "http://www.grossonapoletano.com"},
	{"Calle Gran Via 6", "Madrid", "Spain", "Oven", "http://www.oven.es"},
	{"Halmstadvägen 1", "Falkenberg", "Sweden", "Lilla Napoli", "http://www.lillanapoli.se"},
	{"Gibraltargatan 20", "Goteborg", "Sweden", "Bov", "http://www.bovgbg.com"},
	{"Kopparvägen 30", "Onsala", "Sweden", "Onsala Pizzeria", "http://www.onsalapizzeria.se"},
	{"Boulevard Georges-Favon 12", "Geneva", "Switzerland", "Kytaly", "https://kytaly.ch"},
	{"Chemin de la Tourelle 2", "Geneva", "Switzerland", "Luigia", "http://www.luigia.ch"},
	{"Pavla Tychyny Avenue 1в", "Kyiv", "Ukraine", "Positano", "https://positano.kiev.ua"},
	{"23-25 Gibson Street", "Glasgow", "UK", "La Favorita", "http://lafav.co.uk/glasgow"},
	{"94 Miller Street", "Glasgow", "UK", "Paesano", "https://paesanopizza.co.uk"},
	{"7 Northumberland Avenue", "London", "UK", "50 Kalò di Ciro Salvo", "https://www.50kalo.it/ciro__salvo.php"},
	{"17-21 Sternhold Avenue", "London", "UK", "Addomé", "http://www.addomme.co.uk"},
	{"2 Sir Simon Milton Square", "London", "UK", "Hai Cenato", "http://www.haicenato.co.uk"},
	{"66 Heath Street", "London", "UK", "L'Antica Pizzeria", "http://www.anticapizzeria.co.uk"},
	{"199 Baker Street", "London", "UK", "L'Antica Pizzeria da Michele", "https://www.anticapizzeriadamichele.co.uk"},
	{"182 Bellenden Road", "London", "UK", "Made of Dough", "http://www.madeofdough.co.uk"},
	{"11 Dean Street", "London", "UK", "Pizza Pilgrims", "http://pizzapilgrims.co.uk"},
	{"135 Wardour Street", "London", "UK", "Princi", "http://www.princi.com"},
	{"59-61 Exmouth Market", "London", "UK", "Santoré", "http://www.santorerestaurant.london"},
	{"63 Pitfield Street", "London", "UK", "Sodo Pizza Hoxton", "https://www.sodopizza.co.uk"},
	{"368 Barlow Moor Road", "Manchester", "UK", "Double Zero", "https://www.pizzeriadoublezero.com"},
	{"9 Cotton Street", "Manchester", "UK", "Rudy's Neapolitan Pizza", "https://www.rudyspizza.co.uk"},
	{"38-40 Adelaide Road", "Southampton", "UK", "It's a Pizza Thing", "https://www.facebook.com/apizzathing"},
	{"Rua Graúna 125", "São Paulo", "Brazil", "Braz", "http://www.brazpizzaria.com.br"},
}

// Add a new feature to a collection
func newCollection(c *geogoth.Features, lat float64, lng float64, id string, p *Pizzeria) {
	point := geogoth.NewPoint([]float64{lng, lat})
	feature := geogoth.NewFeature()

	feature.SetID(id)
	feature.SetProperty("city", p.city)
	feature.SetProperty("pizzeria", p.name)
	feature.SetProperty("address", p.address)
	feature.SetProperty("website", p.website)
	feature.SetGeometry(point)
	c.AddFeature(feature)
}

// Save a GeoJSON byte array into a file
func saveGeoJsonToFile(jsondata []byte) {
	f, err := os.Create(GEOJSON_FILE)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err2 := f.Write(jsondata)
	if err2 != nil {
		panic(err2)
	}
}

func main() {
	// Initialize GeoJSON collection
	featureC := geogoth.NewFeatureCollection()
	// Initialize OpenStreetMap geocoder
	g := openstreetmap.Geocoder()
	// Initialize each pizzeria
	for i, n := range pizzaList {
		testAddress := n.address + ", " + n.city
		location, _ := g.Geocode(testAddress)
		if location == nil {
			fmt.Println("got <nil> location")
			fmt.Println("Pizzeria that couldn't be found was ", testAddress)
			continue
		}
		address, _ := g.ReverseGeocode(location.Lat, location.Lng)
		if address != nil {
			fmt.Printf("Address of (%.6F, %.6f) is %s\n", location.Lat, location.Lng, address.FormattedAddress)
		} else {
			continue
		}

		// Save the just found point on map into a new feature of the GeoJSON 
		// collection
		newCollection(featureC, location.Lat, location.Lng, string(i), &n)
	}
	// Have the collection made into GeoJSON
	m, _ := json.Marshal(featureC)
	// Save the final collection to a file
	saveGeoJsonToFile(m)
}
