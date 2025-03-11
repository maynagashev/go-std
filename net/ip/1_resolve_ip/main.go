package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"strings"
)

var (
	bindAddr             = flag.String("bind_addr", ":8080", "")
	resolveIPUsingHeader = flag.Bool("resolve_ip_using_header", false, "")
)

func main() {
	flag.Parse()

	s := &server{bindAddr: *bindAddr, resolveIPUsingHeader: *resolveIPUsingHeader}
	s.ListenAndServe()
}

type server struct {
	bindAddr             string
	resolveIPUsingHeader bool
}

func (s *server) ListenAndServe() {
	http.Handle("/get_ip", http.HandlerFunc(s.handleRequest))
	http.ListenAndServe(s.bindAddr, nil)
}

func (s *server) handleRequest(w http.ResponseWriter, r *http.Request) {
	ip, err := resolveIP(r, resolveIPOpts{UseHeader: s.resolveIPUsingHeader})
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(200)
	w.Write([]byte(ip))
}

type resolveIPOpts struct {
	UseHeader bool
}

func resolveIP(r *http.Request, opts resolveIPOpts) (net.IP, error) {
	if !opts.UseHeader {
		addr := r.RemoteAddr
		// метод возвращает адрес в формате host:port
		// нужна только подстрока host
		ipStr, _, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}
		// парсим ip
		ip := net.ParseIP(ipStr)
		if ip == nil {
			panic("unexpected parse ip error")
		}
		return ip, nil
	} else {
		// смотрим заголовок запроса X-Real-IP
		ipStr := r.Header.Get("X-Real-IP")
		// парсим ip
		ip := net.ParseIP(ipStr)
		if ip == nil {
			// если заголовок X-Real-IP пуст, пробуем X-Forwarded-For
			// этот заголовок содержит адреса отправителя и промежуточных прокси
			// в виде 203.0.113.195, 70.41.3.18, 150.172.238.178
			ips := r.Header.Get("X-Forwarded-For")
			// разделяем цепочку адресов
			ipStrs := strings.Split(ips, ",")
			// интересует только первый
			ipStr = ipStrs[0]
			// парсим
			ip = net.ParseIP(ipStr)
		}
		if ip == nil {
			return nil, fmt.Errorf("failed parse ip from http header")
		}
		return ip, nil
	}
}
