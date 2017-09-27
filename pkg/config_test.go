package pkg

import "testing"

func TestAddEnvironment_WithNoError(t *testing.T) {
	env := NewEnvironment("bogus-env-name", "bogus-domain-name")
	cfg := Config{Default: "path", Domain: "domain", Environments: []Environment{env}}
	err := cfg.AddEnvironment(Environment{Name: "another-bogus-name"})

	if err != nil {
		t.Error("unexpected error:", err)
	}
}

func TestAddEnvironment_WithError(t *testing.T) {
	env := NewEnvironment("bogus-env-name", "bogus-domain-name")
	cfg := Config{Default: "path", Domain: "domain", Environments: []Environment{env}}
	err := cfg.AddEnvironment(env)

	if err == nil {
		t.Error("expected an error got nil")
	}
}

func TestRemoveEnvironment_WithNoError(t *testing.T) {
	env1 := NewEnvironment("bogus-env-name-1", "bogus-domain-name")
	env2 := NewEnvironment("bogus-env-name-2", "bogus-domain-name")
	cfg := Config{Default: env1.Name, Environments: []Environment{env1, env2}}
	err := cfg.RemoveEnvironment(env2.Name)

	if err != nil {
		t.Error("unexpected error:", err)
	}
}

func TestRemoveEnvironment_WithError(t *testing.T) {
	env := NewEnvironment("bogus-env-name-", "bogus-domain-name")
	cfg := Config{Default: env.Name, Environments: []Environment{env}}
	err := cfg.RemoveEnvironment(env.Name)

	if err == nil {
		t.Error("expected an error got nil")
	}
}

func TestLoadEnvironment_WithNoError(t *testing.T) {
	env := NewEnvironment("bogus-env-name-1", "bogus-domain-name")
	cfg := Config{Default: env.Name, Environments: []Environment{env}}
	exp, err := cfg.LoadEnvironment(env.Name)

	if err != nil {
		t.Error("unexpected error:", err)
	}
	if exp.Name != env.Name {
		t.Error("expected env got nil")
	}
}

func TestLoadEnvironment_WithError(t *testing.T) {
	env := NewEnvironment("bogus-env-name-1", "bogus-domain-name")
	cfg := Config{Default: env.Name, Environments: []Environment{env}}
	exp, err := cfg.LoadEnvironment("unknown-env-name")

	if err == nil {
		t.Error("expected an error got nil")
	}
	if exp.Name != "" {
		t.Error("expected nil env, got:", env.Name)
	}
}

func TestSwitchEnvironment_WithNoError(t *testing.T) {
	env1 := NewEnvironment("bogus-env-name-1", "bogus-domain-name")
	env2 := NewEnvironment("bogus-env-name-2", "bogus-domain-name")
	cfg := Config{Default: env1.Name, Environments: []Environment{env1, env2}}
	err := cfg.SwitchEnvironment(env2.Name)

	if err != nil {
		t.Error("unexpected error:", err)
	}
	if cfg.Default != env2.Name {
		t.Errorf("expected %s got %s", env2.Name, cfg.Default)
	}
}

func TestSwitchEnvironment_WithError(t *testing.T) {
	env1 := NewEnvironment("bogus-env-name-1", "bogus-domain-name")
	env2 := NewEnvironment("bogus-env-name-2", "bogus-domain-name")
	cfg := Config{Default: env1.Name, Environments: []Environment{env1}}
	err := cfg.SwitchEnvironment(env2.Name)

	if err == nil {
		t.Error("expected an error got nil")
	}
	if cfg.Default != env1.Name {
		t.Errorf("expected %s got %s", env1.Name, cfg.Default)
	}
}
