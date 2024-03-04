package web

import "github.com/diamondburned/arikawa/v3/discord"

const embedColor = discord.Color(16146719)

func createDJEmbed(req *djApplication) discord.Embed {
	e := discord.Embed{
		Title: "DJ Application",
		Color: embedColor,
	}

	experience := "No"
	if req.Experience {
		experience = "Yes"
	}

	e.Fields = append(e.Fields,
		discord.EmbedField{
			Name:   "DJ Name",
			Value:  req.Name,
			Inline: false,
		},
		discord.EmbedField{
			Name:   "Discord Username",
			Value:  req.Username,
			Inline: false,
		},
		discord.EmbedField{
			Name:   "Preferred Day",
			Value:  req.PreferredDay,
			Inline: false,
		},
		discord.EmbedField{
			Name:   "Experience?",
			Value:  experience,
			Inline: false,
		},
		discord.EmbedField{
			Name:   "Genre(s)",
			Value:  req.Genre,
			Inline: false,
		},
		discord.EmbedField{
			Name:   "Pronouns",
			Value:  req.Pronouns,
			Inline: false,
		},
	)

	return e
}
