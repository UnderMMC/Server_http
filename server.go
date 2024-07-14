package main

import (
	"encoding/json" // Пакет для работы с JSON, который предоставляет функции для кодирования (marshal) и декодирования (unmarshal) данных;
	"net/http"      // Пакет для работы с HTTP-протоколом, включая создание веб-серверов и обработку HTTP-запросов;
	"strconv"
)

type Response struct { // Cтруктурные теги для управления именами полей
	Message string `json:"message"` // или изменения имен при декодировании.
	Name    string `json:"name"`
	Age     uint   `json:"age"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) { // Функция-обработчик HTTP-запросов;
	response := Response{Message: "Hello, Vlad!", Name: "My name is Max, " + " my age", Age: 22}
	w.Header().Set("Content-Type", "application/json") // Устанавливается заголовок ответ,  чтобы клиент мог понять формат ответа;
	json.NewEncoder(w).Encode(response)
}

func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	ageString := r.FormValue("age")
	age, err := strconv.Atoi(ageString)
	if err != nil {
		http.Error(w, "Invalid age value", http.StatusBadRequest)
		return
	}
	result := make(map[string]interface{})
	if age < 18 {
		result["message"] = "Ты слишком молод..."
	} else {
		result["message"] = "Ты слишком стар..."
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/api/hello", HelloHandler) // Регистрирует функцию для обработки HTTP-запросов на эндпоинт /api/hello, через GET-запрос на этот эндпоинт, будет вызван helloHandler;
	http.ListenAndServe(":8080", nil)           // Запускает HTTP-сервер на порту 8080. nil передается как параметр Handler,что означает, что будет использоваться DefaultServeMux (стандартный мультиплексор сервера);
	http.HandleFunc("/api/question", QuestionHandler)
	http.ListenAndServe(":8080", nil)
}
