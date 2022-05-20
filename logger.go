/* Copyright (c) 2020 vesoft inc. All rights reserved.
 *
 * This source code is licensed under Apache 2.0 License,
 * attached with Common Clause Condition 1.0, found in the LICENSES directory.
 */

package nebula_go

import (
	"log"
)

var (
	_ Logger = (*DefaultLogger)(nil)
	_ Logger = (*NoLogger)(nil)
)

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

type DefaultLogger struct{}

func (l DefaultLogger) Info(msg string) {
	log.Printf("[INFO] %s\n", msg)
}

func (l DefaultLogger) Warn(msg string) {
	log.Printf("[WARNING] %s\n", msg)
}

func (l DefaultLogger) Error(msg string) {
	log.Printf("[ERROR] %s\n", msg)
}

func (l DefaultLogger) Fatal(msg string) {
	log.Fatalf("[FATAL] %s\n", msg)
}

// NoLogger type. Will omit all messages except Fatal.
type NoLogger struct{}

// Info method.
func (l NoLogger) Info(msg string) {}

// Warn method.
func (l NoLogger) Warn(msg string) {}

// Error method.
func (l NoLogger) Error(msg string) {}

// Fatal method.
func (l NoLogger) Fatal(msg string) {
	log.Fatalf("[FATAL] %s\n", msg)
}
