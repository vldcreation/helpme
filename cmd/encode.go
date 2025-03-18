package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vldcreation/helpme-package/pkg/encode"
	"github.com/vldcreation/helpme/util"
)

type encodeCmd struct {
	cmd *cobra.Command

	// flags
	encoder         string
	source          string
	format          string
	copyToClipboard bool
	withMimeType    bool
}

func NewEncodeCommand() *encodeCmd {
	apps := &encodeCmd{}
	cmd := &cobra.Command{
		Use:   "encode",
		Short: "encode file or text",
		Long:  "encode file or text",
		Args:  cobra.ExactArgs(0),
	}

	cmd.PersistentFlags().StringVarP(&apps.encoder, "encoder", "e", "", "Source encoder to encode(eg. file | text default: text)")
	cmd.PersistentFlags().StringVarP(&apps.source, "source", "s", "", "Source of file or text to encode (eg. /mypath/myfile.png | helloworld)")
	cmd.PersistentFlags().StringVarP(&apps.format, "format", "f", "", "format encoder to use (eg. base64 | hex). default base64")
	cmd.PersistentFlags().BoolVarP(&apps.copyToClipboard, "clipboard", "c", false, "Copy to clipboard")
	cmd.PersistentFlags().BoolVarP(&apps.withMimeType, "mimetype", "m", false, "Show mime type")

	cmd.MarkPersistentFlagRequired("source")

	apps.cmd = cmd
	return apps
}

func (c *encodeCmd) Command() *cobra.Command {
	c.cmd.Run = c.Execute
	return c.cmd
}

func (c *encodeCmd) Execute(_ *cobra.Command, args []string) {
	encoder := switchEncoder(c.source, c.encoder)
	applyFormatEncoder(encoder, c.source, c.format)
	encoder.ApplyOpt(encode.WithCopyToClipboard(c.copyToClipboard), encode.WithMimeType(c.withMimeType))

	out, err := encoder.Encode()
	if err != nil {
		println(err.Error())
		return
	}

	util.PrintlnGreen(out)
}

func switchEncoder(source string, encoder string) encode.Encoder {
	switch encoder {
	case "file":
		return encode.NewFileEncoder(source)
	case "text":
		return encode.NewTextEncoder(source)
	default:
		return encode.NewTextEncoder(source)
	}
}

func applyFormatEncoder(e encode.Encoder, source string, format string) {
	switch format {
	case "base64":
		e.ApplyOpt(encode.WithFormatEncoder(encode.NewBase64Encoder(source)))
	case "base32":
		e.ApplyOpt(encode.WithFormatEncoder(encode.NewBase32Encoder(source)))
	default:
		e.ApplyOpt(encode.WithFormatEncoder(encode.NewBase64Encoder(source)))
	}
}
