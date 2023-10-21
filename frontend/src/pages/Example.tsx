import {useState} from 'react';
import './App.css';
import {Greet} from "../../wailsjs/go/main/App";
import {Button} from "@/components/ui/button";
import * as z from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {useForm} from "react-hook-form";
import {Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {ModeToggle} from "@/components/mode-toggle";
import {Client} from "@/util/client";
import {useMutation, useQuery, useQueryClient} from "@tanstack/react-query";

const formSchema = z.object({
  username: z.string().min(2, {
    message: "Username must be at least 2 characters.",
  }),
})

function Example() {
  const [resultText, setResultText] = useState("Please enter your name below 👇");
  const [name, setName] = useState('');
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    Greet(name).then(updateResultText);
  }

  // Access the client
  const queryClient = useQueryClient()

  // Queries
  const query = useQuery({ queryKey: ['example'], queryFn: ()=> {return Client().dev.getDevExample()} })

  // Mutations
  const mutation = useMutation({
    mutationFn: Client().okr.postOkrObjective,
    onSuccess: () => {
      // Invalidate and refetch
      queryClient.invalidateQueries({ queryKey: ['example'] })
    },
  })

  // 1. Define your form.
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
    },
  })

  // 2. Define a submit handler.
  function onSubmit(values: z.infer<typeof formSchema>) {
    // Do something with the form values.
    // ✅ This will be type-safe and validated.
    console.log(values)

    Client().dev.getDevExample().then(res => console.log(res));
  }

  return (
    <>
      <div id="result" className="result">{resultText}</div>
      <ModeToggle/>
      <div id="input" className="input-box">
        <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
        <Button onClick={greet}>Click me</Button>
      </div>
      <div>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
            <FormField
              control={form.control}
              name="username"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Username</FormLabel>
                  <FormControl>
                    <Input placeholder="shadcn" {...field} />
                  </FormControl>
                  <FormDescription>
                    This is your public display name.
                  </FormDescription>
                  <FormMessage/>
                </FormItem>
              )}
            />
            <Button type="submit">Submit</Button>
          </form>
        </Form>
      </div>
      <div>
        title: {query.data?.data.title}
      </div>
    </>
  )
}

export default Example
