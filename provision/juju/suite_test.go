// Copyright 2013 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package juju

import (
	"github.com/globocom/config"
	"github.com/globocom/tsuru/db"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct {
	collName string
	conn     *db.Storage
}

var _ = Suite(&S{})

func (s *S) SetUpSuite(c *C) {
	s.collName = "juju_units_test"
	config.Set("git:host", "tsuruhost.com")
	config.Set("juju:units-collection", s.collName)
	err := handler.DryRun()
	c.Assert(err, IsNil)
	config.Set("database:url", "127.0.0.1:27017")
	config.Set("database:name", "juju_provision_tests_s")
	s.conn, err = db.Conn()
	c.Assert(err, IsNil)
}

func (s *S) TearDownSuite(c *C) {
	handler.Stop()
	handler.Wait()
	s.conn.Collection(s.collName).Database.DropDatabase()
}
