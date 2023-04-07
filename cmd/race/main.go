package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/tetra3d"
)

type Game struct {
	GameScene *tetra3d.Scene
	Camera    *tetra3d.Camera
}

func NewGame() *Game {

	g := &Game{}

	// First, we load a scene from a .gltf or .glb file. LoadGLTFFile takes a filepath and
	// any loading options (nil can be taken as a valid default set of loading options), and
	// returns a *tetra3d.Library and an error if it was unsuccessful. We can also use
	// tetra3d.LoadGLTFData() if we don't have access to the host OS's filesystem (if the
	// assets are embedded, for example).

	//
	//library, err := tetra3d.LoadGLTFFile("example.gltf", nil)
	//library, err := tetra3d.LoadGLTFFile("assets/evo_horn.gltf", nil)
	//library, err := tetra3d.LoadGLTFFile("assets/horn.gltf", nil)
	//library, err := tetra3d.LoadGLTFFile("assets/next.gltf", nil)

	//library, err := tetra3d.LoadGLTFFile("assets/next_toaster.gltf", nil)
	library, err := tetra3d.LoadGLTFFile("assets/arma.gltf", nil)

	if err != nil {
		panic(err)
	}

	// A Library is essentially everything that got exported from your 3D modeler -
	// all of the scenes, meshes, materials, and animations. The ExportedScene of a Library
	// is the scene that was active when the file was exported.

	// We'll clone the ExportedScene so we don't change it irreversibly; making a clone
	// of a Tetra3D resource (Scene, Node, Material, Mesh, Camera, whatever) makes a deep
	// copy of it.
	g.GameScene = library.ExportedScene.Clone()

	// Tetra3D uses OpenGL's coordinate system (+X = Right, +Y = Up, +Z = Backward [towards the camera]),
	// in comparison to Blender's coordinate system (+X = Right, +Y = Forward,
	// +Z = Up). Note that when loading models in via GLTF or DAE, models are
	// converted automatically (so up is +Z in Blender and +Y in Tetra3D automatically).

	// We could create a new Camera as below - we would pass the size of the screen to the
	// Camera so it can create its own buffer textures (which are *ebiten.Images).

	// g.Camera = tetra3d.NewCamera(ScreenWidth, ScreenHeight)

	// However, we can also just grab an existing camera from the scene if it
	// were exported from the GLTF file - if exported through Blender's Tetra3D add-on,
	// then the camera size can be easily set from within Blender.

	g.Camera = g.GameScene.Root.Get("Camera").(*tetra3d.Camera)

	// Camera implements the tetra3d.INode interface, which means it can be placed
	// in 3D space and can be parented to another Node somewhere in the scene tree.
	// Models, Lights, and Nodes (which are essentially "empties" one can
	// use for positioning and parenting) can, as well.

	// We can place Models, Cameras, and other Nodes with node.SetWorldPosition() or
	// node.SetLocalPosition(). There are also variants that take a 3D Vector.

	// The *World variants of positioning functions takes into account absolute space;
	// the Local variants position Nodes relative to their parents' positioning and
	// transforms (and is more performant.)
	// You can also move Nodes using Node.Move(x, y, z) / Node.MoveVec(vector).

	// Each Scene has a tree that starts with the Root Node. To add Nodes to the Scene,
	// parent them to the Scene's base, like so:

	// scene.Root.AddChildren(object)

	// To remove them, you can unparent them from either the parent (Node.RemoveChildren())
	// or the child (Node.Unparent()). When a Node is unparented, it is removed from the scene
	// tree; if you want to destroy the Node, then dropping any references to this Node
	// at this point would be sufficient.

	// For Cameras, we don't actually need to place them in a scene to view the Scene, since
	// the presence of the Camera in the Scene node tree doesn't impact what it would see.

	// We can see the tree "visually" by printing out the hierarchy:
	fmt.Println(g.GameScene.Root.HierarchyAsString())

	// You can also visualize the scene hierarchy using TetraTerm:
	// https://github.com/SolarLune/tetraterm

	return g
}

func (g *Game) Update() error { return nil }

func (g *Game) Draw(screen *ebiten.Image) {

	// Here, we'll call Camera.Clear() to clear its internal backing texture. This
	// should be called once per frame before drawing your Scene.
	g.Camera.Clear()

	// Now we'll render the Scene from the camera. The Camera's ColorTexture will then
	// hold the result.

	// Camera.RenderScene() renders all Nodes in a scene, starting with the
	// scene's root. You can also use Camera.Render() to simply render a selection of
	// individual Models, or Camera.RenderNodes() to render a subset of a scene tree.
	g.Camera.RenderScene(g.GameScene)

	// To see the result, we draw the Camera's ColorTexture to the screen.
	// Before doing so, we'll clear the screen first. In this case, we'll do this
	// with a color, though we can also go with screen.Clear().
	screen.Fill(color.RGBA{20, 30, 40, 255})

	// Draw the resulting texture to the screen, and you're done! You can
	// also visualize the depth texture with g.Camera.DepthTexture().
	screen.DrawImage(g.Camera.ColorTexture(), nil)

	// Note that the resulting texture is indeed just an ordinary *ebiten.Image, so
	// you can also use this as a texture for a Model's Material, as an example.

}

func (g *Game) Layout(w, h int) (int, int) {

	// Here, by simply returning the camera's size, we are essentially
	// scaling the camera's output to the window size and letterboxing as necessary.

	// If you wanted to extend the camera according to window size, you would
	// have to resize the camera using the window's new width and height.
	return g.Camera.Size()

}

func main() {

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

}
