"use client"


import React, { useState, useEffect } from "react";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useCookies } from "next-client-cookies";

export const TodoSelect = () => {

  const [placeholder, setPlaceholder] = useState<string>("All todos");
  

  return (
    <Select>
      <SelectTrigger className="w-56">
        <SelectValue placeholder={placeholder} />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectItem value="all tasks">
            All tasks
          </SelectItem>
          <SelectItem value="pending">
            Pending
          </SelectItem>
          <SelectItem value="pending">
            Completed
          </SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};