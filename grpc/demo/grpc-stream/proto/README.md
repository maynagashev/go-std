# Стриминг с сервера

Рассмотрим пример потокового отправления данных со стороны gRPC-сервера. 

Предположим, что клиент отправляет серверу целое число и максимальный множитель, 
а сервер с интервалом в 100 миллисекунд умножает это число на 1, 2, 3 
и далее до максимального множителя и возвращает результаты клиенту. 