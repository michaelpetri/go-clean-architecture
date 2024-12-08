package application

import (
	"com.michael-petri/todo/adapter/cli/commands"
	"com.michael-petri/todo/application/usecase"
	"com.michael-petri/todo/domain/model"
	"com.michael-petri/todo/domain/repository"
	"com.michael-petri/todo/infrastructure/in_memory"
	"context"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewCliApplication() *fx.App {
	return fx.New(
		// Register repository
		fx.Provide(func() repository.TodoRepository {
			state := make(map[int64]*model.Todo)
			return in_memory.NewInMemoryTodoRepository(&state)
		}),
		// Register use-cases
		fx.Provide(func(todos repository.TodoRepository) usecase.CreateTodoCase {
			return *usecase.NewCreateTodoCase(todos)
		}),
		fx.Provide(func(todos repository.TodoRepository) usecase.ListTodosCase {
			return *usecase.NewListTodosCase(todos)
		}),
		fx.Provide(func(todos repository.TodoRepository) usecase.ResolveTodoCase {
			return *usecase.NewResolveTodoCase(todos)
		}),
		// Register cobra commands
		fx.Provide(
			func(
				createCase usecase.CreateTodoCase,
				listCase usecase.ListTodosCase,
				resolveCase usecase.ResolveTodoCase,
			) *cobra.Command {
				root := &cobra.Command{
					Use:   "todo",
					Short: "A simple todo tool",
					Long:  `A simple todo tool which can create, list and resolve todos.`,
				}

				root.AddCommand(
					commands.NewCreateTodoCommand(&createCase),
					commands.NewListTodosCommand(&listCase),
					commands.NewResolveTodoCommand(&resolveCase),
					commands.NewInteractiveCommand(
						&createCase,
						&listCase,
						&resolveCase,
					),
				)

				return root
			},
		),
		// Register custom lifecycle to execute cobra root command
		fx.Invoke(func(lifecycle fx.Lifecycle, root *cobra.Command, shutdowner fx.Shutdowner) {
			lifecycle.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						exitCode := 0
						if err := root.Execute(); err != nil {
							exitCode = 1
						}
						_ = shutdowner.Shutdown(fx.ExitCode(exitCode))
					}()
					return nil
				},
			})
		}),
	)
}
