// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package what_changed

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestCompareDiscriminator_PropertyNameChanged(t *testing.T) {

	left := `propertyName: chicken`

	right := `propertyName: nuggets`

	var lNode, rNode yaml.Node
	_ = yaml.Unmarshal([]byte(left), &lNode)
	_ = yaml.Unmarshal([]byte(right), &rNode)

	// create low level objects
	var lDoc base.Discriminator
	var rDoc base.Discriminator
	_ = low.BuildModel(&lNode, &lDoc)
	_ = low.BuildModel(&rNode, &rDoc)

	// compare.
	extChanges := CompareDiscriminator(&lDoc, &rDoc)
	assert.Equal(t, 1, extChanges.TotalChanges())
	assert.Equal(t, Modified, extChanges.Changes[0].ChangeType)

}

func TestCompareDiscriminator_PropertyNameRemoved(t *testing.T) {

	left := `propertyName: chicken`

	right := ``

	var lNode, rNode yaml.Node
	_ = yaml.Unmarshal([]byte(left), &lNode)
	_ = yaml.Unmarshal([]byte(right), &rNode)

	// create low level objects
	var lDoc base.Discriminator
	var rDoc base.Discriminator
	_ = low.BuildModel(&lNode, &lDoc)
	_ = low.BuildModel(&rNode, &rDoc)

	// compare.
	extChanges := CompareDiscriminator(&lDoc, &rDoc)
	assert.Equal(t, 1, extChanges.TotalChanges())
	assert.Equal(t, PropertyRemoved, extChanges.Changes[0].ChangeType)
}

func TestCompareDiscriminator_PropertyNameAdded(t *testing.T) {

	left := ``

	right := `propertyName: chicken`

	var lNode, rNode yaml.Node
	_ = yaml.Unmarshal([]byte(left), &lNode)
	_ = yaml.Unmarshal([]byte(right), &rNode)

	// create low level objects
	var lDoc base.Discriminator
	var rDoc base.Discriminator
	_ = low.BuildModel(&lNode, &lDoc)
	_ = low.BuildModel(&rNode, &rDoc)

	// compare.
	extChanges := CompareDiscriminator(&lDoc, &rDoc)
	assert.Equal(t, 1, extChanges.TotalChanges())
	assert.Equal(t, PropertyAdded, extChanges.Changes[0].ChangeType)
}

func TestCompareDiscriminator_MappingAdded(t *testing.T) {

	left := `propertyName: chicken`

	right := `propertyName: chicken
mapping:
  chuffing: puffing
  hacking: coding`

	var lNode, rNode yaml.Node
	_ = yaml.Unmarshal([]byte(left), &lNode)
	_ = yaml.Unmarshal([]byte(right), &rNode)

	// create low level objects
	var lDoc base.Discriminator
	var rDoc base.Discriminator
	_ = low.BuildModel(&lNode, &lDoc)
	_ = low.BuildModel(&rNode, &rDoc)

	// compare.
	extChanges := CompareDiscriminator(&lDoc, &rDoc)
	assert.Equal(t, 2, extChanges.TotalChanges())
	assert.Equal(t, ObjectAdded, extChanges.MappingChanges[0].ChangeType)
	assert.Equal(t, ObjectAdded, extChanges.MappingChanges[1].ChangeType)
	assert.Equal(t, "chuffing", extChanges.MappingChanges[0].Property)
	assert.Equal(t, "puffing", extChanges.MappingChanges[0].New)
	assert.Equal(t, "hacking", extChanges.MappingChanges[1].Property)
	assert.Equal(t, "coding", extChanges.MappingChanges[1].New)

}

func TestCompareDiscriminator_MappingRemoved(t *testing.T) {

	left := `propertyName: chicken
mapping:
  chuffing: puffing
  hacking: coding`

	right := `propertyName: chicken
mapping:
  hacking: coding`

	var lNode, rNode yaml.Node
	_ = yaml.Unmarshal([]byte(left), &lNode)
	_ = yaml.Unmarshal([]byte(right), &rNode)

	// create low level objects
	var lDoc base.Discriminator
	var rDoc base.Discriminator
	_ = low.BuildModel(&lNode, &lDoc)
	_ = low.BuildModel(&rNode, &rDoc)

	// compare.
	extChanges := CompareDiscriminator(&lDoc, &rDoc)
	assert.Equal(t, 1, extChanges.TotalChanges())
	assert.Equal(t, ObjectRemoved, extChanges.MappingChanges[0].ChangeType)
	assert.Equal(t, "chuffing", extChanges.MappingChanges[0].Property)
	assert.Equal(t, "puffing", extChanges.MappingChanges[0].Original)
}

func TestCompareDiscriminator_SingleMappingAdded(t *testing.T) {

	left := `propertyName: chicken
mapping:
  chuffing: puffing`

	right := `propertyName: chicken
mapping:
  chuffing: puffing
  hacking: coding`

	var lNode, rNode yaml.Node
	_ = yaml.Unmarshal([]byte(left), &lNode)
	_ = yaml.Unmarshal([]byte(right), &rNode)

	// create low level objects
	var lDoc base.Discriminator
	var rDoc base.Discriminator
	_ = low.BuildModel(&lNode, &lDoc)
	_ = low.BuildModel(&rNode, &rDoc)

	// compare.
	extChanges := CompareDiscriminator(&lDoc, &rDoc)
	assert.Equal(t, 1, extChanges.TotalChanges())
	assert.Equal(t, ObjectAdded, extChanges.MappingChanges[0].ChangeType)
	assert.Equal(t, "hacking", extChanges.MappingChanges[0].Property)
	assert.Equal(t, "coding", extChanges.MappingChanges[0].New)

}

func TestCompareDiscriminator_MultiMappingAdded(t *testing.T) {

	left := `propertyName: chicken
mapping:
  chuffing: puffing`

	right := `propertyName: chicken
mapping:
  chuffing: puffing
  hacking: coding
  singing: dancing`

	var lNode, rNode yaml.Node
	_ = yaml.Unmarshal([]byte(left), &lNode)
	_ = yaml.Unmarshal([]byte(right), &rNode)

	// create low level objects
	var lDoc base.Discriminator
	var rDoc base.Discriminator
	_ = low.BuildModel(&lNode, &lDoc)
	_ = low.BuildModel(&rNode, &rDoc)

	// compare.
	extChanges := CompareDiscriminator(&lDoc, &rDoc)
	assert.Equal(t, 2, extChanges.TotalChanges())
	assert.Equal(t, ObjectAdded, extChanges.MappingChanges[0].ChangeType)
	assert.Equal(t, "hacking", extChanges.MappingChanges[0].Property)
	assert.Equal(t, "coding", extChanges.MappingChanges[0].New)
	assert.Equal(t, "singing", extChanges.MappingChanges[1].Property)
	assert.Equal(t, "dancing", extChanges.MappingChanges[1].New)

}

func TestCompareDiscriminator_SingleMappingModified(t *testing.T) {

	left := `propertyName: chicken
mapping:
  chuffing: puffing`

	right := `propertyName: chicken
mapping:
  chuffing: herbs`

	var lNode, rNode yaml.Node
	_ = yaml.Unmarshal([]byte(left), &lNode)
	_ = yaml.Unmarshal([]byte(right), &rNode)

	// create low level objects
	var lDoc base.Discriminator
	var rDoc base.Discriminator
	_ = low.BuildModel(&lNode, &lDoc)
	_ = low.BuildModel(&rNode, &rDoc)

	// compare.
	extChanges := CompareDiscriminator(&lDoc, &rDoc)
	assert.Equal(t, 1, extChanges.TotalChanges())
	assert.Equal(t, Modified, extChanges.MappingChanges[0].ChangeType)
	assert.Equal(t, "chuffing", extChanges.MappingChanges[0].Property)
	assert.Equal(t, "herbs", extChanges.MappingChanges[0].New)
	assert.Equal(t, "puffing", extChanges.MappingChanges[0].Original)

	// should be a single breaking change
	assert.Equal(t, 1, CountBreakingChanges(extChanges.MappingChanges))

}

func TestCompareDiscriminator_Identical(t *testing.T) {

	left := `propertyName: chicken`

	right := `propertyName: chicken`

	var lNode, rNode yaml.Node
	_ = yaml.Unmarshal([]byte(left), &lNode)
	_ = yaml.Unmarshal([]byte(right), &rNode)

	// create low level objects
	var lDoc base.Discriminator
	var rDoc base.Discriminator
	_ = low.BuildModel(&lNode, &lDoc)
	_ = low.BuildModel(&rNode, &rDoc)

	// compare.
	extChanges := CompareDiscriminator(&lDoc, &rDoc)
	assert.Nil(t, extChanges)
}