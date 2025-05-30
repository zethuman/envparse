package envparse

import (
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"strings"
	"time"
)

const logMsg = "using default value %s=%v"
const logDefInvalid = "default value is invalid %s=%v"

type parser struct {
	name  string
	value string
	def   any
}

func Env(name string) *parser {
	return &parser{
		name:  name,
		value: os.Getenv(name),
	}
}

func (p *parser) Def(def any) *parser {
	p.def = def
	return p
}

func (p *parser) Str() string {
	def, ok := p.def.(string)
	if !ok {
		log.Warn().Msgf(logDefInvalid, p.name, p.def)
	}

	if p.value == "" {
		log.Warn().Msgf(logMsg, p.name, def)
		return def
	}
	return p.value
}

func (p *parser) StrArr() []string {
	def, ok := p.def.([]string)
	if !ok {
		log.Warn().Msgf(logDefInvalid, p.name, p.def)
	}

	if p.value == "" {
		log.Warn().Msgf(logMsg, p.name, def)
		return def
	}
	return strings.Split(p.value, ",")
}

func (p *parser) Int() int {
	def, ok := p.def.(int)
	if !ok {
		log.Warn().Msgf(logDefInvalid, p.name, p.def)
	}

	if p.value == "" {
		log.Warn().Msgf(logMsg, p.name, def)
		return def
	}

	parsed, err := strconv.Atoi(p.value)
	if err != nil {
		log.Warn().Err(err).Msgf(logMsg, p.name, def)
		return def
	}

	return parsed
}

func (p *parser) Duration() time.Duration {
	def, ok := p.def.(time.Duration)
	if !ok {
		log.Warn().Msgf(logDefInvalid, p.name, p.def)
	}

	if p.value == "" {
		log.Warn().Msgf(logMsg, p.name, def)
		return def
	}

	parsed, err := time.ParseDuration(p.value)
	if err != nil {
		log.Warn().Err(err).Msgf(logMsg, p.name, def)
		return def
	}

	return parsed
}

func (p *parser) Bool() bool {
	def, ok := p.def.(bool)
	if !ok {
		log.Warn().Msgf(logDefInvalid, p.name, p.def)
	}

	if p.value == "" {
		log.Warn().Msgf(logMsg, p.name, def)
		return def
	}

	parsed, err := strconv.ParseBool(p.value)
	if err != nil {
		log.Warn().Err(err).Msgf(logMsg, p.name, def)
		return def
	}

	return parsed
}

func (p *parser) BoolPtr() *bool {
	def, ok := p.def.(bool)
	if !ok {
		log.Warn().Msgf(logDefInvalid, p.name, p.def)
	}

	if p.value == "" {
		log.Warn().Msgf(logMsg, p.name, def)
		return &def
	}

	parsed, err := strconv.ParseBool(p.value)
	if err != nil {
		log.Warn().Err(err).Msgf(logMsg, p.name, def)
		return &def
	}

	return &parsed
}
