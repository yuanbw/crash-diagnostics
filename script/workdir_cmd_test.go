// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package script

import (
	"fmt"
	"os"
	"testing"
)

func TestCommandWORKDIR(t *testing.T) {
	tests := []commandTest{
		{
			name: "WORKDIR with default param",
			source: func() string {
				return "WORKDIR foo/bar"
			},
			script: func(s *Script) error {
				dirs := s.Preambles[CmdWorkDir]
				if len(dirs) != 1 {
					return fmt.Errorf("Script has unexpected number of WORKDIR %d", len(dirs))
				}
				wdCmd, ok := dirs[0].(*WorkdirCommand)
				if !ok {
					return fmt.Errorf("Unexpected type %T in script", dirs[0])
				}
				if wdCmd.Path() != "foo/bar" {
					return fmt.Errorf("WORKDIR has unexpected directory %s", wdCmd.Path())
				}
				return nil
			},
		},
		{
			name: "Multiple WORKDIRs",
			source: func() string {
				return "WORKDIR foo/bar\nWORKDIR 'bazz/buzz'"
			},
			script: func(s *Script) error {
				dirs := s.Preambles[CmdWorkDir]
				if len(dirs) != 1 {
					return fmt.Errorf("Script has unexpected number of WORKDIR %d", len(dirs))
				}
				wdCmd, ok := dirs[0].(*WorkdirCommand)
				if !ok {
					return fmt.Errorf("Unexpected type %T in script", dirs[0])
				}
				if wdCmd.Path() != "bazz/buzz" {
					return fmt.Errorf("WORKDIR has unexpected directory %s", wdCmd.Path())
				}
				return nil
			},
		},
		{
			name: "WORKDIR with named param",
			source: func() string {
				return "WORKDIR path:foo/bar"
			},
			script: func(s *Script) error {
				dirs := s.Preambles[CmdWorkDir]
				if len(dirs) != 1 {
					return fmt.Errorf("Script has unexpected number of WORKDIR %d", len(dirs))
				}
				wdCmd, ok := dirs[0].(*WorkdirCommand)
				if !ok {
					return fmt.Errorf("Unexpected type %T in script", dirs[0])
				}
				if wdCmd.Path() != "foo/bar" {
					return fmt.Errorf("WORKDIR has unexpected directory %s", wdCmd.Path())
				}
				return nil
			},
		},
		{
			name: "WORKDIR with quoted named param",
			source: func() string {
				return "WORKDIR path:'foo/bar'"
			},
			script: func(s *Script) error {
				dirs := s.Preambles[CmdWorkDir]
				if len(dirs) != 1 {
					return fmt.Errorf("Script has unexpected number of WORKDIR %d", len(dirs))
				}
				wdCmd, ok := dirs[0].(*WorkdirCommand)
				if !ok {
					return fmt.Errorf("Unexpected type %T in script", dirs[0])
				}
				if wdCmd.Path() != "foo/bar" {
					return fmt.Errorf("WORKDIR has unexpected directory %s", wdCmd.Path())
				}
				return nil
			},
		},
		{
			name: "WORKDIR with expanded vars",
			source: func() string {
				os.Setenv("foopath", "foo/bar")
				return "WORKDIR path:'${foopath}'"
			},
			script: func(s *Script) error {
				dirs := s.Preambles[CmdWorkDir]
				if len(dirs) != 1 {
					return fmt.Errorf("Script has unexpected number of WORKDIR %d", len(dirs))
				}
				wdCmd, ok := dirs[0].(*WorkdirCommand)
				if !ok {
					return fmt.Errorf("Unexpected type %T in script", dirs[0])
				}
				if wdCmd.Path() != "foo/bar" {
					return fmt.Errorf("WORKDIR has unexpected directory %s", wdCmd.Path())
				}
				return nil
			},
		},
		{
			name: "WORKDIR with multiple args",
			source: func() string {
				return "WORKDIR foo/bar bazz/buzz"
			},
			shouldFail: true,
		},
		{
			name: "WORKDIR with no args",
			source: func() string {
				return "WORKDIR"
			},
			shouldFail: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			runCommandTest(t, test)
		})
	}
}
