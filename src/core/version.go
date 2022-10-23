package core

import (
	"ahkpm/src/utils"
	"errors"
	"strings"
)

type Version interface {
	VersionKind() VersionKind
	Value() string
	String() string
}

type version struct {
	kind  VersionKind
	value string
}

type VersionKind string

const (
	SemVerExact VersionKind = "Semantic Version"
	Branch      VersionKind = "Branch"
	Tag         VersionKind = "Tag"
	Commit      VersionKind = "Commit"
)

func NewVersion(kind VersionKind, value string) Version {
	return version{
		kind:  VersionKind(kind),
		value: value,
	}
}

// Converts a version specifier string into a Version.
func VersionFromSpecifier(versionSpecifier string) (Version, error) {
	v := version{}

	if utils.IsSemVer(versionSpecifier) {
		v.kind = SemVerExact
		v.value = versionSpecifier
	} else if strings.HasPrefix(versionSpecifier, "branch:") {
		v.kind = Branch
		v.value = strings.TrimPrefix(versionSpecifier, "branch:")
	} else if strings.HasPrefix(versionSpecifier, "tag:") {
		v.kind = Tag
		v.value = strings.TrimPrefix(versionSpecifier, "tag:")
	} else if strings.HasPrefix(versionSpecifier, "commit:") {
		v.kind = Commit
		v.value = strings.TrimPrefix(versionSpecifier, "commit:")
	} else {
		return v, errors.New("Invalid version specifier " + versionSpecifier)
	}

	return v, nil
}

// Represents the Version as a valid version specifier string.
func (v version) String() string {
	if v.kind == Branch || v.kind == Tag || v.kind == Commit {
		return strings.ToLower(string(v.kind)) + ":" + v.value
	}
	if v.kind == SemVerExact {
		return v.value
	}
	utils.Exit("Invalid version kind " + string(v.kind))
	return ""
}

func (v version) VersionKind() VersionKind {
	return v.kind
}

func (v version) Value() string {
	return v.value
}
