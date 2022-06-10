package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type JsonDate time.Time

type Response struct {
	FechaReporteWeb       JsonDate `json:"fecha_reporte_web"`
	IdDeCaso              int      `json:"id_de_caso,string"`
	FechaDeNotificacion   JsonDate `json:"fecha_de_notificaci_n"`
	Departamento          int      `json:"departamento,string"`
	DepartamentoNombre    string   `json:"departamento_nom"`
	CiudadMunicipio       int      `json:"ciudad_municipio,string"`
	CiudadMunicipioNombre string   `json:"ciudad_municipio_nom"`
	Edad                  int      `json:"edad,string"`
	UnidadMedida          int      `json:"unidad_medida,string"`
	Sexo                  string   `json:"sexo"`
	FuenteTipoContagio    string   `json:"fuente_tipo_contagio"`
	Ubicacion             string   `json:"ubicacion"`
	Estado                string   `json:"estado"`
	Pais                  int      `json:"pais,string"`
	PaisNombre            string   `json:"pais_nom"`
	Recuperado            string   `json:"recuperado"`
	FechaInicioSintomas   JsonDate `json:"fecha_inicio_sintomas"`
	FechaMuerte           JsonDate `json:"fecha_muerte"`
	FechaDiagnostico      JsonDate `json:"fecha_diagnostico"`
	FechaRecuperacion     JsonDate `json:"fecha_recuperado"`
	TipoRecuperacion      string   `json:"tipo_recuperacion"`
	PertenenciaEtnica     int      `json:"per_etn_,string"`
	NombreGrupoEtnico     string   `json:"nom_grupo_"`
}

func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return nil
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	resp, err := http.Get("https://www.datos.gov.co/resource/gt2j-8ykr.json")
	if err != nil {
		fmt.Println("No response from request")
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body")
			log.Fatal(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		log.Fatal(err)
	}
	var result []Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot unmarshal JSON")
		log.Fatal(err)
	}
	fmt.Println(PrettyPrint(result))
}
