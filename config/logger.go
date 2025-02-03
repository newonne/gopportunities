package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := os.Stdout
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	//o operador & cria um ponteiro para a instância da estrutura Logger. Isso significa
	// que a função NewLogger retorna um ponteiro para o objeto recém-criado, em vez de retornar uma cópia.
	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// create non-formatted logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// create format enabled logs
func (l *Logger) Debugf(fomart string, v ...interface{}) {
	l.debug.Printf(fomart, v...)
}

func (l *Logger) Infof(fomart string, v ...interface{}) {
	l.info.Printf(fomart, v...)
}

func (l *Logger) Warningf(fomart string, v ...interface{}) {
	l.warning.Printf(fomart, v...)
}

func (l *Logger) Errorf(fomart string, v ...interface{}) {
	l.err.Printf(fomart, v...)
}
