package profiler

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/pprof"
	"time"
)

// Profiler handles CPU and memory profiling
type Profiler struct {
	cpuFile   *os.File
	memFile   *os.File
	outputDir string
}

// New creates a new profiler instance
func New(outputDir string) (*Profiler, error) {
	if outputDir == "" {
		outputDir = "."
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	return &Profiler{outputDir: outputDir}, nil
}

// StartCPUProfile starts CPU profiling
func (p *Profiler) StartCPUProfile() error {
	timestamp := time.Now().Format("20060102-150405")
	cpuFile := filepath.Join(p.outputDir, fmt.Sprintf("cpu-%s.prof", timestamp))

	f, err := os.Create(cpuFile)
	if err != nil {
		return fmt.Errorf("could not create CPU profile: %w", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		f.Close()
		return fmt.Errorf("could not start CPU profile: %w", err)
	}

	p.cpuFile = f
	return nil
}

// StopCPUProfile stops CPU profiling
func (p *Profiler) StopCPUProfile() {
	if p.cpuFile != nil {
		pprof.StopCPUProfile()
		p.cpuFile.Close()
		p.cpuFile = nil
	}
}

// WriteHeapProfile writes memory profile
func (p *Profiler) WriteHeapProfile() error {
	timestamp := time.Now().Format("20060102-150405")
	memFile := filepath.Join(p.outputDir, fmt.Sprintf("mem-%s.prof", timestamp))

	f, err := os.Create(memFile)
	if err != nil {
		return fmt.Errorf("could not create memory profile: %w", err)
	}
	defer f.Close()

	if err := pprof.WriteHeapProfile(f); err != nil {
		return fmt.Errorf("could not write memory profile: %w", err)
	}

	return nil
}
