package config

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

type FileDecorator struct {
	Icon  string       `yaml:"icon"`
	Hex   string       `yaml:"hex"`
	Color *color.Color `yaml:"-"`
}

func (fd *FileDecorator) UnmarshalYAML(value *yaml.Node) error {
	type rawDecorator FileDecorator
	var raw rawDecorator
	if err := value.Decode(&raw); err != nil {
		return err
	}

	r, g, b, err := hexToRgb(raw.Hex)
	if err != nil {
		return err
	}
	raw.Color = color.RGB(r, g, b)
	*fd = FileDecorator(raw)
	return nil
}

func hexToRgb(hex string) (int, int, int, error) {
	var r, g, b int
	_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	if err != nil {
		return 0, 0, 0, err
	}
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return 0, 0, 0, fmt.Errorf("invalid RGB value")
	}
	return r, g, b, nil
}

var DefaultFileDecorators = map[string]FileDecorator{
	".txt":  {Icon: "📝", Hex: "#808080", Color: color.RGB(128, 128, 128)},
	".md":   {Icon: "📝", Hex: "#808080", Color: color.RGB(128, 128, 128)},
	".pdf":  {Icon: "📕", Hex: "#FF4C4C", Color: color.RGB(255, 76, 76)},
	".docx": {Icon: "📘", Hex: "#2B579A", Color: color.RGB(43, 87, 154)},
	".xlsx": {Icon: "📗", Hex: "#217346", Color: color.RGB(33, 115, 70)},
	".pptx": {Icon: "📙", Hex: "#D24726", Color: color.RGB(210, 71, 38)},
	".csv":  {Icon: "📄", Hex: "#2E8B57", Color: color.RGB(46, 139, 87)},
	".json": {Icon: "🧾", Hex: "#F1C40F", Color: color.RGB(241, 196, 15)},
	".yml":  {Icon: "🧾", Hex: "#F1C40F", Color: color.RGB(241, 196, 15)},
	".yaml": {Icon: "🧾", Hex: "#F1C40F", Color: color.RGB(241, 196, 15)},
	".xml":  {Icon: "🧾", Hex: "#8E44AD", Color: color.RGB(142, 68, 173)},
	".html": {Icon: "🌐", Hex: "#E67E22", Color: color.RGB(230, 126, 34)},
	".css":  {Icon: "🎨", Hex: "#3498DB", Color: color.RGB(52, 152, 219)},
	".scss": {Icon: "🎨", Hex: "#3498DB", Color: color.RGB(52, 152, 219)},
	".less": {Icon: "🎨", Hex: "#3498DB", Color: color.RGB(52, 152, 219)},
	".sass": {Icon: "🎨", Hex: "#3498DB", Color: color.RGB(52, 152, 219)},
	".js":   {Icon: "⚙️", Hex: "#F7DF1E", Color: color.RGB(247, 223, 30)},
	".ts":   {Icon: "⚙️", Hex: "#3178C6", Color: color.RGB(49, 120, 198)},
	".go":   {Icon: "🐹", Hex: "#00ADD8", Color: color.RGB(0, 173, 216)},
	".py":   {Icon: "🐍", Hex: "#3572A5", Color: color.RGB(53, 114, 165)},
	".java": {Icon: "☕", Hex: "#B07219", Color: color.RGB(176, 114, 25)},
	".c":    {Icon: "💻", Hex: "#555555", Color: color.RGB(85, 85, 85)},
	".h":    {Icon: "💻", Hex: "#AAAAAA", Color: color.RGB(170, 170, 170)},
	".cpp":  {Icon: "💻", Hex: "#F34B7D", Color: color.RGB(243, 75, 125)},
	".hpp":  {Icon: "💻", Hex: "#B0C4DE", Color: color.RGB(176, 196, 222)},
	".rb":   {Icon: "💎", Hex: "#701516", Color: color.RGB(112, 21, 22)},
	".php":  {Icon: "🐘", Hex: "#4F5D95", Color: color.RGB(79, 93, 149)},
	".sh":   {Icon: "🖥️", Hex: "#586E75", Color: color.RGB(88, 110, 117)},
	".bat":  {Icon: "🖥️", Hex: "#586E75", Color: color.RGB(88, 110, 117)},
	".exe":  {Icon: "⚙️", Hex: "#000000", Color: color.RGB(0, 0, 0)},
	".zip":  {Icon: "📦", Hex: "#95A5A6", Color: color.RGB(149, 165, 166)},
	".tar":  {Icon: "📦", Hex: "#95A5A6", Color: color.RGB(149, 165, 166)},
	".gz":   {Icon: "📦", Hex: "#95A5A6", Color: color.RGB(149, 165, 166)},
	".bz2":  {Icon: "📦", Hex: "#95A5A6", Color: color.RGB(149, 165, 166)},
	".xz":   {Icon: "📦", Hex: "#95A5A6", Color: color.RGB(149, 165, 166)},
	".7z":   {Icon: "📦", Hex: "#95A5A6", Color: color.RGB(149, 165, 166)},
	".rar":  {Icon: "📦", Hex: "#95A5A6", Color: color.RGB(149, 165, 166)},
	".jpg":  {Icon: "🖼️", Hex: "#9B59B6", Color: color.RGB(155, 89, 182)},
	".jpeg": {Icon: "🖼️", Hex: "#9B59B6", Color: color.RGB(155, 89, 182)},
	".png":  {Icon: "🖼️", Hex: "#9B59B6", Color: color.RGB(155, 89, 182)},
	".gif":  {Icon: "🖼️", Hex: "#9B59B6", Color: color.RGB(155, 89, 182)},
	".svg":  {Icon: "🖼️", Hex: "#9B59B6", Color: color.RGB(155, 89, 182)},
	".bmp":  {Icon: "🖼️", Hex: "#9B59B6", Color: color.RGB(155, 89, 182)},
	".tiff": {Icon: "🖼️", Hex: "#9B59B6", Color: color.RGB(155, 89, 182)},
	".webp": {Icon: "🖼️", Hex: "#9B59B6", Color: color.RGB(155, 89, 182)},
	".mp3":  {Icon: "🎵", Hex: "#E67E22", Color: color.RGB(230, 126, 34)},
	".wav":  {Icon: "🎵", Hex: "#E67E22", Color: color.RGB(230, 126, 34)},
	".mp4":  {Icon: "🎬", Hex: "#C0392B", Color: color.RGB(192, 57, 43)},
	".avi":  {Icon: "🎬", Hex: "#C0392B", Color: color.RGB(192, 57, 43)},
	".mkv":  {Icon: "🎬", Hex: "#C0392B", Color: color.RGB(192, 57, 43)},
	".mov":  {Icon: "🎬", Hex: "#C0392B", Color: color.RGB(192, 57, 43)},
	".iso":  {Icon: "💿", Hex: "#34495E", Color: color.RGB(52, 73, 94)},
}

var DefaultFolderDecorator = FileDecorator{
	Icon:  "📁",
	Hex:   "#2980B9",
	Color: color.RGB(41, 128, 185),
}
