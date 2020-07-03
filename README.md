# mod4

```
mod2->mod1(v1.0.0)
mod3->mod1(v1.0.1)
mod4->mod1(v1.0.2)

mod1.Version(): mod1{ mod_version:1.0.2 go:go1.14.3 }

mod2.Version(): mod1{ mod_version:1.0.2 go:go1.14.3 }
mod2{ mod_version:1.0.0 go:go1.14.3 }

mod3.Version(): mod1{ mod_version:1.0.2 go:go1.14.3 }
mod3{ mod_version:1.0.1 go:go1.14.3 }
```