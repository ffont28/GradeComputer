package main
import "fmt"
func main() {
  var max int
  fmt.Println("-------------")
  fmt.Println("Version PC_1.1.1")
  fmt.Println("No copyright claimed by the author ffont28. For reports and/or suggestions contact on telegram @ffontana")
  fmt.Println("----")
  fmt.Println("Benvenuto nel programma di calcolo di voto in proporzione.")
  fmt.Println("Per uscire clicca sulla X qui sopra nel terminale.")
  fmt.Print("Inserisci il massimo del punteggio che corrisponde al 10: ")
  fmt.Scan(&max) 
  for max<100000{
    var punteggio int
        fmt.Print("Inserisci il punteggio ottenuto su ", max, ": ");
        fmt.Scan(&punteggio);
    var votoFinale float64 = (float64(punteggio)*10.0)/float64(max)
        fmt.Printf("VOTO: %.2f   ||   VOTO CON CRITERIO ITIS: %.2f ", votoFinale, votoFinale-0.5);
    fmt.Println()
        fmt.Println("-------------------------------------------------");
  }
}
