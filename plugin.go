package exdocs

import (
	"fmt"
	"strings"

	"github.com/vito/booklit"
)

func init() {
	booklit.RegisterPlugin("ex-docs", NewPlugin)
}

type Plugin struct {
	section *booklit.Section

	definitionContext []string
	noteIdx           int
}

func NewPlugin(section *booklit.Section) booklit.Plugin {
	return &Plugin{
		section: section,
	}
}

func (p Plugin) Ghuser(user string) booklit.Content {
	return booklit.Styled{
		Style:   "github-user-link",
		Content: booklit.String(user),
	}
}

func (p Plugin) Ghrelease(tag string, optionalRepo ...string) booklit.Content {
	repo := "concourse"
	if len(optionalRepo) > 0 {
		repo = optionalRepo[0]
	}

	return booklit.Styled{
		Style:   "github-release-link",
		Content: booklit.String(tag),
		Partials: booklit.Partials{
			"Owner": booklit.String("concourse"),
			"Repo":  booklit.String(repo),
		},
	}
}

func (p Plugin) Ghpr(number string, optionalRepo ...string) booklit.Content {
	repo := "concourse"
	if len(optionalRepo) > 0 {
		repo = optionalRepo[0]
	}

	return booklit.Styled{
		Style:   "github-pr-link",
		Content: booklit.String(number),
		Partials: booklit.Partials{
			"Owner": booklit.String("concourse"),
			"Repo":  booklit.String(repo),
		},
	}
}

func (p Plugin) Ghissue(number string, optionalRepo ...string) booklit.Content {
	repo := "concourse"
	if len(optionalRepo) > 0 {
		repo = optionalRepo[0]
	}

	return booklit.Styled{
		Style:   "github-issue-link",
		Content: booklit.String(number),
		Partials: booklit.Partials{
			"Owner": booklit.String("concourse"),
			"Repo":  booklit.String(repo),
		},
	}
}

func (p Plugin) PipelineImage(path string) booklit.Content {
	return booklit.Styled{
		Style: "pipeline-image",
		Content: booklit.Image{
			Path:        path,
			Description: "pipeline",
		},
	}
}

func (p *Plugin) Note(commaSeparatedTags string, content booklit.Content) booklit.Content {
	tags := strings.Split(commaSeparatedTags, ",")

	p.noteIdx++

	idx := p.noteIdx
	targetTag := fmt.Sprintf("%s-note-%d", p.section.PrimaryTag.Name, idx)

	tagNotes := []booklit.Content{}
	for _, t := range tags {
		tagNotes = append(tagNotes, booklit.Styled{
			Style:   "release-note-tag",
			Content: booklit.String(t),
		})
	}

	return booklit.Styled{
		Style:   "release-note",
		Content: content,
		Partials: booklit.Partials{
			"Tags":   booklit.Sequence(tagNotes),
			"Target": booklit.String(targetTag),
		},
	}
}
