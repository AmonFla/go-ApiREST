package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Declaro el paquete de datos de pruebas
var Sample map[string]Personas

//Creo la estructura  de Personas, utilizado para manipular los datos
type Personas struct {
	Id       string
	Nombre   string
	Apellido string
}

//Inicializo los datos de ejemplo
func init() {
	Sample = map[string]Personas{
		"1": Personas{Id: "1", Nombre: "Luis", Apellido: "Perez"},
		"2": Personas{Id: "2", Nombre: "Maria", Apellido: "Romano"},
		"3": Personas{Id: "3", Nombre: "Nestor", Apellido: "Sanchez"},
	}
}

//Función para obtener todos los datos de las personas
func GetPersonas(w http.ResponseWriter, r *http.Request) {
	//Codifico el map en json
	salidaJson, _ := json.Marshal(Sample)

	//Inidico que los datos devueltos van en json
	w.Header().Set("Content-Type", "application/json")
	w.Write(salidaJson)
}

//Retorno un solo valor.
func GetPersonaById(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	//Verifico si existe y retorno el valor en json o doy un error NotFound
	if _, ok := Sample[vars["ID"]]; ok {
		//Codifico el registro en json
		salidaJson, _ := json.Marshal(Sample[vars["ID"]])
		//Inidico que los datos devueltos van en json
		w.Header().Set("Content-Type", "application/json")
		w.Write(salidaJson)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

//Guardo nuevos valores
func SavePersona(w http.ResponseWriter, r *http.Request) {
	//Decodifico dentro de persona el body del request
	var persona Personas
	err := json.NewDecoder(r.Body).Decode(&persona)

	//Valido si se pudo decodificar
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//Agrego el nuevo valor a Sample.
	Sample[persona.Id] = persona

	// Retorno la persona generada
	salidaJson, _ := json.Marshal(Sample[persona.Id])
	// Indico que los datos devueltos van en json
	w.Header().Set("Content-Type", "application/json")
	w.Write(salidaJson)

}

//Actualizo un registro
func EditPersona(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	//Verifico si existe y actualizo o doy un error NotFound
	if _, ok := Sample[vars["ID"]]; ok {
		//Decodifico dentro de persona el body del request
		var persona Personas
		err := json.NewDecoder(r.Body).Decode(&persona)
		//Valido si se pudo decodificar
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			//Actualizo el nuevo valor a Sample.
			Sample[vars["ID"]] = persona
			// Retorno la persona modificada
			salidaJson, _ := json.Marshal(Sample[vars["ID"]])
			// Indico que los datos devueltos van en json
			w.Header().Set("Content-Type", "application/json")
			w.Write(salidaJson)

		}
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

}

//Elimino el registro indicado
func DelPersona(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	//Verifico si existe y eliino o doy un error NotFound
	if _, ok := Sample[vars["ID"]]; ok {
		delete(Sample, vars["ID"])
		http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
