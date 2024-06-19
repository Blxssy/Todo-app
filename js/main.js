document.addEventListener('DOMContentLoaded', () => {
    const todoList = document.getElementById('todo-list');
    const todoForm = document.getElementById('todo-form');
    const todoTitle = document.getElementById('todo-title');

    const fetchTodos = () => {
        fetch('/todos')
            .then(response => response.json())
            .then(todos => {
                todoList.innerHTML = '';
                todos.forEach(todo => {
                    const listItem = createTodoListItem(todo);
                    todoList.appendChild(listItem);
                });
            })
            .catch(error => console.error('Error fetching todos:', error));
    };

    const addTodo = (title) => {
        fetch('/todos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ title })
        })
        .then(response => response.json())
        .then(newTodo => {
            const listItem = createTodoListItem(newTodo);
            todoList.appendChild(listItem);
            todoTitle.value = '';
        })
        .catch(error => console.error('Error adding todo:', error));
    };

    const completeTodo = (id) => {
        fetch(`/todos/${id}/complete`, {
            method: 'PUT'
        })
        .then(response => response.json())
        .then(updatedTodo => {
            const listItem = document.getElementById(`todo-${updatedTodo.id}`);
            if (listItem) {
                listItem.classList.add('completed');
            }
        })
        .catch(error => console.error('Error completing todo:', error));
    };

    const createTodoListItem = (todo) => {
        const listItem = document.createElement('li');
        listItem.id = `todo-${todo.id}`;
        listItem.textContent = todo.title;
        if (todo.state) {
            listItem.classList.add('completed');
        }

        const completeButton = document.createElement('button');
        completeButton.textContent = 'Complete';
        completeButton.addEventListener('click', () => completeTodo(todo.id));

        listItem.appendChild(completeButton);
        return listItem;
    };

    todoForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const title = todoTitle.value.trim();
        if (title) {
            addTodo(title);
        }
        fetchTodos();
    });

    fetchTodos();
});
