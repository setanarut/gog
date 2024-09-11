[![GoDoc](https://godoc.org/github.com/setanarut/gog/v2?status.svg)](https://pkg.go.dev/github.com/setanarut/gog/v2)

# GOG v2

GOG is a Go Object-oriented 2d drawing library for creative coding and generative art

![curve_anim](https://github.com/user-attachments/assets/135882e0-5a6d-438c-b0d0-80ebd12713c2)

## Examples

### Path follow animation

```Go
ctx := gog.NewContext(250, 100)
rect := shapes.Rect(vec.Vec2{}, 30, 10)
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
```

![path_follow](https://github.com/user-attachments/assets/55ac6887-41eb-4fb1-8d1c-55e2cdcb93fb)


See folder [examples](./examples) for all examples
