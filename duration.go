package config

import "time"

type DurationValidator func(time.Duration) error

type durationValidators struct {
	*durationValue
	validators []DurationValidator
}

func (v durationValidators) Set(s string) error {
	err := v.durationValue.Set(s)
	if err != nil {
		return err
	}

	for _, validator := range v.validators {
		err = validator(time.Duration(*v.durationValue))
		if err != nil {
			return err
		}
	}

	return nil
}

type durationValue time.Duration

func (d durationValue) String() string {
	return time.Duration(d).String()
}

func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = durationValue(v)
	return err
}

func Duration(v *time.Duration, flag, env, usage string, validators ...DurationValidator) *Flag {
	return &Flag{
		Value: durationValidators{
			durationValue: (*durationValue)(v),
			validators:    validators,
		},
		Name:     flag,
		Env:      env,
		Usage:    usage,
		TypeHint: "duration",
	}
}