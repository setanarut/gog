# Path follow animation

```Go
func main() {
	ctx := gog.NewContext(250, 100)
	rect := shapes.Rect(v.Vec{}, 30, 10)
	lemn := shapes.Lemniscate(100, 100).SetPos(ctx.Center)
	lemnTotalLength := lemn.Length()
	for _, length := range utils.Linspace(0, lemnTotalLength, 120) {
		ctx.Clear(color.Black)
		ctx.Stroke(lemn, gog.DefaultStrokeStyle())
		pos, ang := lemn.PointAngleAtLength(length)
		ctx.Fill(rect.SetPos(pos).Rotated(ang), color.White)
		ctx.AppendAnimationFrame()
	}
	ctx.SaveAPNG("path_follow.png", 2)
}
```

![path_follow](https://github.com/user-attachments/assets/55ac6887-41eb-4fb1-8d1c-55e2cdcb93fb)
