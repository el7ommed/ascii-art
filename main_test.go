package main

import (
	"os/exec"
	"testing"
)

func Test001(t *testing.T) {
	want := ""

	got, _ := exec.Command("go", "run", "main.go", "").Output()
	if want != string(got) {
		t.Errorf("Expected:\n%s \nGot:\n%s ", want, string(got))
	}

}

func Test002(t *testing.T) {
	want := "\n"

	got, _ := exec.Command("go", "run", "main.go", "\\n").Output()
	if want != string(got) {
		t.Errorf("Expected:\n%s \nGot:\n%s ", want, string(got))
	}

}

func Test003(t *testing.T) {
	want := " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n\n"

	got, _ := exec.Command("go", "run", "main.go", "Hello\\n").Output()
	if want != string(got) {
		t.Errorf("Expected:\n%s \nGot:\n%s ", want, string(got))
	}

}

func Test004(t *testing.T) {
	want := " _              _   _          \n| |            | | | |         \n| |__     ___  | | | |   ___   \n|  _ \\   / _ \\ | | | |  / _ \\  \n| | | | |  __/ | | | | | (_) | \n|_| |_|  \\___| |_| |_|  \\___/  \n                               \n                               \n"

	got, _ := exec.Command("go", "run", "main.go", "hello").Output()
	if want != string(got) {
		t.Errorf("Expected:\n%s \nGot:\n%s ", want, string(got))
	}

}

func Test005(t *testing.T) {
	want := " _    _          _        _    ____   \n| |  | |        | |      | |  / __ \\  \n| |__| |   ___  | |      | | | |  | | \n|  __  |  / _ \\ | |      | | | |  | | \n| |  | | |  __/ | |____  | | | |__| | \n|_|  |_|  \\___| |______| |_|  \\____/  \n                                      \n                                      \n"

	got, _ := exec.Command("go", "run", "main.go", "HeLlO").Output()
	if want != string(got) {
		t.Errorf("Expected:\n%s \nGot:\n%s ", want, string(got))
	}

}

func Test006(t *testing.T) {
	want := " _    _          _   _                 _______   _                           \n| |  | |        | | | |               |__   __| | |                          \n| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  \n|  __  |  / _ \\ | | | |  / _ \\           | |    |  _ \\   / _ \\ | '__|  / _ \\ \n| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ \n|_|  |_|  \\___| |_| |_|  \\___/           |_|    |_| |_|  \\___| |_|     \\___| \n                                                                             \n                                                                             \n"

	got, _ := exec.Command("go", "run", "main.go", "Hello There").Output()
	if want != string(got) {
		t.Errorf("Expected:\n%s \nGot:\n%s ", want, string(got))
	}

}

func Test007(t *testing.T) {
	want := "     _    _          _   _                         _______   _                           \n _  | |  | |        | | | |                ____   |__   __| | |                          \n/ | | |__| |   ___  | | | |   ___         |___ \\     | |    | |__     ___   _ __    ___  \n| | |  __  |  / _ \\ | | | |  / _ \\          __) |    | |    |  _ \\   / _ \\ | '__|  / _ \\ \n| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ \n|_| |_|  |_|  \\___| |_| |_|  \\___/        |_____|    |_|    |_| |_|  \\___| |_|     \\___| \n                                                                                         \n                                                                                         \n"

	got, _ := exec.Command("go", "run", "main.go", "1Hello 2There").Output()
	if want != string(got) {
		t.Errorf("Expected:\n%s \nGot:\n%s ", want, string(got))
	}

}

func Test008(t *testing.T) {
	want := "   __  _    _          _   _                 _______   _                           __    \n  / / | |  | |        | | | |               |__   __| | |                          \\ \\   \n | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  \n/ /   |  __  |  / _ \\ | | | |  / _ \\           | |    |  _ \\   / _ \\ | '__|  / _ \\   \\ \\ \n\\ \\   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / \n | |  |_|  |_|  \\___| |_| |_|  \\___/           |_|    |_| |_|  \\___| |_|     \\___|  | |  \n  \\_\\                                                                              /_/   \n                                                                                         \n"

	got, _ := exec.Command("go", "run", "main.go", "{Hello There}").Output()
	if want != string(got) {
		t.Errorf("Expected:\n%s \nGot:\n%s ", want, string(got))
	}

}

func Test009(t *testing.T) {
	want := " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n _______   _                           \n|__   __| | |                          \n   | |    | |__     ___   _ __    ___  \n   | |    |  _ \\   / _ \\ | '__|  / _ \\ \n   | |    | | | | |  __/ | |    |  __/ \n   |_|    |_| |_|  \\___| |_|     \\___| \n                                       \n                                       \n"

	got, _ := exec.Command("go", "run", "main.go", "Hello\\nThere").Output()
	if want != string(got) {
		t.Errorf("Expected:\n%s \nGot:\n%s ", want, string(got))
	}

}

func Test010(t *testing.T) {
	want := " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n\n _______   _                           \n|__   __| | |                          \n   | |    | |__     ___   _ __    ___  \n   | |    |  _ \\   / _ \\ | '__|  / _ \\ \n   | |    | | | | |  __/ | |    |  __/ \n   |_|    |_| |_|  \\___| |_|     \\___| \n                                       \n                                       \n"

	got, _ := exec.Command("go", "run", "main.go", "Hello\\n\\nThere").Output()
	if want != string(got) {
		t.Errorf("Expected:\n%s \nGot:\n%s ", want, string(got))
	}

}
