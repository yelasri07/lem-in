package functions

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func (a *Info) ReadFile(arg string) error {
	file, err := os.Open(arg)
	if err != nil {
		return errors.New("error to open the file")
	}

	sc := bufio.NewScanner(file)

	na := false
	start := false
	s := 0
	end := false
	e := 0

	for sc.Scan() {

		l := sc.Text()
		if l == "" {
			continue
		}

		line := strings.TrimSpace(l)

		if !na {
			na = true

			var err error

			a.NumberOfAnts, err = strconv.Atoi(line)
			if err != nil {
				return errors.New("invalid nm ants")
			}
			if a.NumberOfAnts <= 0 {
				return errors.New("invalid nm ants")
			}
		}

		if line == "##start" {
			start = true
			s++
			continue
		}

		if start && s == 1 {
			nmofstart, status := validestratorend(line)
			if status {
				a.Start = nmofstart
				start = false

			} else {
				continue
			}

		}

		if line == "##end" {
			end = true
			e++
			continue
		}
		if end && e == 1 {
			nmofend, status := validestratorend(line)

			if status {
				a.End = nmofend
				end = false

			} else {
				continue
			}
		}

		if s > 1 || e > 1 {
			return errors.New("invalid form of the file")
		}

		if line[0] != '#' {
			err := a.FindRooms(line)
			if err != nil {
				return errors.New("invalid form of the file")
			}
		}
	}

	if s == 1 && a.Start == "" {
		return errors.New("invalid form of the file")
	}
	if e == 1 && a.End == "" {
		return errors.New("invalid form of the file")
	}

	return nil
}

func (a *Info) FindRooms(line string) error {
	if len(strings.Fields(line)) == 3 {
		sl := strings.Fields(line)
		if !vali_coordinates(sl[1], sl[2]) {
			return errors.New("invalid form of the file")
		}
		if a.Rooms == nil {
			a.Rooms = make(map[string][]string)
		}
		a.Rooms[sl[0]] = append(a.Rooms[sl[0]], sl[1])
		a.Rooms[sl[0]] = append(a.Rooms[sl[0]], sl[2])
		
	}

	if strings.Contains(line, "-") {
		sl := strings.Split(line, "-")
		if len(sl) == 2 {
			if sl[0] == sl[1] {
				return errors.New("invalid form of the file -->tunnels<--")
			}

			if !(a.valid_rooms(sl[0], sl[1])) {
				return errors.New("invalid form of the file jjj")
			}
			if a.Tunnels == nil {
				a.Tunnels = make(map[string][]string)
			}
			if a.isValidToApp(sl[0], sl[1]) {
				a.Tunnels[sl[0]] = append(a.Tunnels[sl[0]], sl[1])
			}
			if a.isValidToApp(sl[1], sl[0]) {
				a.Tunnels[sl[1]] = append(a.Tunnels[sl[1]], sl[0])
			}
		}
	}

	return nil
}

func (y *Info) isValidToApp(r1, r2 string) bool {
	for _, t := range y.Tunnels[r1] {
		if t == r2 {
			return false
		}
	}
	return true
}

func (a *Info) valid_rooms(r1, r2 string) bool {
	if a.Rooms == nil {
		return false
	}
	v1 := false
	v2 := false
	for room := range a.Rooms {
		if r1 == room {
			v1 = true
		}
		if r2 == room {
			v2 = true
		}
	}

	if v1 && v2 {
		return true
	}
	return false
}

func vali_coordinates(c1, c2 string) bool {
	_, err := strconv.Atoi(c1)
	if err != nil {
		return false
	}
	_, errr := strconv.Atoi(c2)
	
	return errr == nil
}

func validestratorend(room string) (string, bool) {
	sl := strings.Fields(room)

	if len(sl) == 3 {
		return sl[0], true
	}
	return "", false
}
