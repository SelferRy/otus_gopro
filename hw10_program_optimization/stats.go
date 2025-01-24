package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

//easyjson:json
type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

var user User

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

//easyjson:json
type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	scanner := bufio.NewScanner(r)

	i := 0
	for scanner.Scan() {
		if err = user.UnmarshalJSON(scanner.Bytes()); err != nil {
			slog.Warn("get user.UnmarshalJSON(scanner.Bytes()). Check error.", slog.Any("error: %w", err))
			return
		}
		result[i] = user
		i++
	}

	if err = scanner.Err(); err != nil {
		slog.Warn("get scanner.Err(). Check error.", slog.Any("error: %w", err))
		return
	}

	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		matched := strings.Contains(user.Email, domain)

		if matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
